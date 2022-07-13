package handlers

import (
	"net/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

func Error(logger zerolog.Logger, translator ut.Translator) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		code := fiber.StatusInternalServerError

		logger.Error().
			Err(err).
			Msg("An error has occurred in application")

		if e, ok := err.(*fiber.Error); ok {
			return ctx.Status(e.Code).JSON(ErrorResponse{
				Message: e.Message,
			})
		}

		if err == ErrInvalidPayload {
			return ctx.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Message: ErrInvalidPayload.Error(),
			})
		}

		if _, ok := err.(*validator.InvalidValidationError); ok {
			return ctx.Status(fiber.StatusUnprocessableEntity).
				JSON(ErrorResponse{Message: "Data is invalid"})
		}

		if err, ok := err.(validator.ValidationErrors); ok {
			return ctx.Status(fiber.StatusUnprocessableEntity).
				JSON(err.Translate(translator))
		}

		return ctx.Status(code).
			JSON(ErrorResponse{Message: "An error has occurred!"})
	}
}
