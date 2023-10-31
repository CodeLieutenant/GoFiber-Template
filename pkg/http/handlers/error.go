package handlers

import (
	"errors"

	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/invopop/validation"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Error(logger zerolog.Logger) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

		if errors.Is(err, primitive.ErrInvalidHex) {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Message: "Invalid JSON Payload, check your input",
			})
		}

		var fiberErr *fiber.Error

		if errors.As(err, &fiberErr) {
			return c.Status(fiberErr.Code).JSON(dto.ErrorResponse{
				Message: fiberErr.Message,
			})
		}

		{
			var validationErr validation.Errors
			if errors.As(err, &validationErr) {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"errors": validationErr,
				})
			}
		}

		logger.Error().Err(err).
			Str("path", c.Route().Path).
			Msg("Failed to process request")

		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: "An error has occurred!",
		})
	}
}
