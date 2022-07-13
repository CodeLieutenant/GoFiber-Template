package testing_utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"

	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/constants"
)

func getBody[T any](headers http.Header, body T) io.Reader {
	switch headers.Get(fiber.HeaderContentType) {
	case fiber.MIMEApplicationJSON, fiber.MIMEApplicationJSONCharsetUTF8:
		jsonStr, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}

		return bytes.NewReader(jsonStr)
	default:
		return nil
	}
}


type RequestModifier func(*http.Request) *http.Request

func WithHeaders(headers http.Header) RequestModifier {
	return func(req *http.Request) *http.Request {
		if headers.Get(fiber.HeaderContentType) == "" {
			headers.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		}

		if headers.Get(fiber.HeaderAccept) == "" {
			headers.Set(fiber.HeaderAccept, fiber.MIMEApplicationJSONCharsetUTF8)
		}

		if headers.Get(fiber.HeaderUserAgent) == "" {
			headers.Set(fiber.HeaderUserAgent, constants.TestUserAgent)
		}

		req.Header = headers

		return req
	}
}

func WithBody[T any](body T) RequestModifier {
	return func(req *http.Request) *http.Request {
		newReq := httptest.NewRequest(req.Method, req.URL.Path, getBody(req.Header, body))
		newReq.Header = req.Header
		newReq.URL.RawQuery = req.URL.Query().Encode()

		for _, cookie := range req.Cookies() {
			newReq.AddCookie(cookie)
		}

		return newReq
	}
}

func WithCookies(cookies []*http.Cookie) RequestModifier {
	return func(req *http.Request) *http.Request {
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}

		return req
	}
}



func MakeRequest[T any](method, uri string, modifiers ...RequestModifier) *http.Request {
	var defaults []func(*http.Request) *http.Request

	switch method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		defaults = []func(*http.Request) *http.Request{
			WithHeaders(http.Header{}),
			WithBody[any](nil),
		}
	default:
		defaults = []func(*http.Request) *http.Request{
			WithHeaders(http.Header{}),
		}
	}

	req := httptest.NewRequest(method, uri, nil)

	for _, modifier := range defaults {
		req = modifier(req)
	}

	for _, modifier := range modifiers {
		req = modifier(req)
	}

	return req
}

func Get(app *fiber.App, uri string, modifiers ...RequestModifier) *http.Response {
	req := MakeRequest[any](http.MethodGet, uri, modifiers...)

	res, err := app.Test(req, -1)
	if err != nil {
		panic("Cannot get response")
	}

	return res
}

func Post[T any](app *fiber.App, uri string, modifiers ...RequestModifier) *http.Response {
	req := MakeRequest[T](http.MethodPost, uri, modifiers...)

	res, err := app.Test(req, -1)
	if err != nil {
		panic("Cannot get response")
	}

	return res
}

func Put[T any](app *fiber.App, uri string, modifiers ...RequestModifier) *http.Response {
	req := MakeRequest[T](http.MethodPut, uri, modifiers...)

	res, err := app.Test(req, -1)
	if err != nil {
		panic("Cannot get response")
	}

	return res
}

func Patch[T any](app *fiber.App, uri string, modifiers ...RequestModifier) *http.Response {
	req := MakeRequest[T](http.MethodPatch, uri, modifiers...)

	res, err := app.Test(req, -1)
	if err != nil {
		panic("Cannot get response")
	}

	return res
}

func Delete(app *fiber.App, uri string, modifiers ...RequestModifier) *http.Response {
	req := MakeRequest[any](http.MethodDelete, uri, modifiers...)

	res, err := app.Test(req, -1)
	if err != nil {
		panic("Cannot get response")
	}

	return res
}
