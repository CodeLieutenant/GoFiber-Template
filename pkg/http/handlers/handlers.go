package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)


var ErrInvalidPayload = errors.New("Invalid Payload")

type ErrorResponse struct {
	Message interface{} `json:"message,omitempty"`
}

const pageNotFoundMessage = "Page is not found"

func NotFound() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		accept := ctx.Get(fiber.HeaderAccept, fiber.MIMEApplicationJSONCharsetUTF8)

		switch accept {
		case fiber.MIMEApplicationJSONCharsetUTF8, fiber.MIMEApplicationJSON:
			return ctx.
				Status(fiber.StatusNotFound).
				JSON(ErrorResponse{Message: pageNotFoundMessage})
		}

		return ctx.SendString(pageNotFoundMessage)
	}
}
