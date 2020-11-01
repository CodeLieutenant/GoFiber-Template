package handlers

import (
	"errors"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type message struct {
	Message string `json:"message"`
}

func Error(translator ut.Translator) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		code := fiber.StatusInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			return ctx.Status(e.Code).JSON(message{
				Message: e.Message,
			})
		}

		if _, ok := err.(*validator.InvalidValidationError); ok {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(message{Message: "Data is invalid"})
		}

		if err, ok := err.(validator.ValidationErrors); ok {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err.Translate(translator))
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(message{Message: "Data not found!"})
		}

		return ctx.Status(code).JSON(message{Message: "An error has occurred!"})
	}
}
