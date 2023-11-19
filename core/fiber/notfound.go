package fiber

import (
	gofiber "github.com/gofiber/fiber/v2"
)

const pageNotFoundMessage = "Page is not found"

func NotFound() gofiber.Handler {
	return func(ctx *gofiber.Ctx) error {
		accept := ctx.Get(gofiber.HeaderAccept, gofiber.MIMEApplicationJSONCharsetUTF8)

		switch accept {
		case gofiber.MIMEApplicationJSONCharsetUTF8, gofiber.MIMEApplicationJSON:
			return ctx.
				Status(gofiber.StatusNotFound).
				JSON(ErrorResponse{Message: pageNotFoundMessage})
		}

		return ctx.SendString(pageNotFoundMessage)
	}
}
