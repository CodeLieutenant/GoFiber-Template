package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupErrorHandlerApp() (*fiber.App, *validator.Validate) {
	v := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	englishTranslations, _ := uni.GetTranslator("en")
	app := fiber.New(fiber.Config{
		ErrorHandler: Error(&log.Logger, englishTranslations),
	})
	return app, v
}

func TestErrorHandler(t *testing.T) {
	t.Parallel()
	asserts := require.New(t)

	t.Run("ReturnFiberError", func(t *testing.T) {
		app, _ := setupErrorHandlerApp()
		app.Get("/", func(ctx *fiber.Ctx) error {
			return fiber.ErrBadGateway
		})
		m := message{}
		res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))

		asserts.Nil(err)
		asserts.EqualValues(fiber.StatusBadGateway, res.StatusCode)
		asserts.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
		asserts.Nil(json.NewDecoder(res.Body).Decode(&m))
		asserts.NotEmpty(m.Message)
	})

	t.Run("ValidationError", func(t *testing.T) {
		app, _ := setupErrorHandlerApp()
		app.Get("/", func(ctx *fiber.Ctx) error {
			return validator.ValidationErrors{}
		})
		res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
		asserts.Nil(err)
		asserts.EqualValues(fiber.StatusUnprocessableEntity, res.StatusCode)
		asserts.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
	})

	t.Run("ValidationError", func(t *testing.T) {
		app, _ := setupErrorHandlerApp()
		app.Get("/", func(ctx *fiber.Ctx) error {
			return &validator.InvalidValidationError{}
		})
		res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
		asserts.Nil(err)
		asserts.EqualValues(fiber.StatusUnprocessableEntity, res.StatusCode)
		asserts.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
	})

	t.Run("NotFound", func(t *testing.T) {
		app, _ := setupErrorHandlerApp()
		app.Get("/", func(ctx *fiber.Ctx) error {
			return gorm.ErrRecordNotFound
		})
		res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
		asserts.Nil(err)
		asserts.EqualValues(fiber.StatusNotFound, res.StatusCode)
		asserts.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
	})
	t.Run("AnyOtherError", func(t *testing.T) {
		app, _ := setupErrorHandlerApp()
		app.Get("/", func(ctx *fiber.Ctx) error {
			return errors.New("any other error")
		})
		res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
		asserts.Nil(err)
		asserts.EqualValues(fiber.StatusInternalServerError, res.StatusCode)
		asserts.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
	})
}
