package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rzajac/zltest"
	"github.com/stretchr/testify/require"

	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/http/handlers"
	"github.com/BrosSquad/GoFiber-Boilerplate/testing_utils"
)

func setupErrorHandlerApp(t *testing.T) (*fiber.App, *validator.Validate, *zltest.Tester) {
	v, translations := testing_utils.GetValidator()

	logger, loggerTest := testing_utils.NewTest(t, zerolog.InfoLevel)

	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.Error(logger, translations),
	})

	return app, v, loggerTest
}

func TestErrorHandler_ReturnFiberError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app, _, _ := setupErrorHandlerApp(t)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return fiber.ErrBadGateway
	})

	m := struct {
		Message string `json:"message"`
	}{}
	res := testing_utils.Get(app, "/")

	assert.EqualValues(fiber.StatusBadGateway, res.StatusCode)
	assert.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
	assert.Nil(json.NewDecoder(res.Body).Decode(&m))
	assert.NotEmpty(m.Message)
}

func TestErrorHandler_InvalidPayloadError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app, _, _ := setupErrorHandlerApp(t)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return handlers.ErrInvalidPayload
	})

	res := testing_utils.Get(app, "/")

	m := struct {
		Message string `json:"message"`
	}{}

	assert.EqualValues(fiber.StatusBadRequest, res.StatusCode)
	assert.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
	assert.Nil(json.NewDecoder(res.Body).Decode(&m))
	assert.NotEmpty(m.Message)
	assert.Equal(handlers.ErrInvalidPayload.Error(), m.Message)
}

func TestErrorHandler_ValidationError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app, _, _ := setupErrorHandlerApp(t)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return validator.ValidationErrors{}
	})
	res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.Nil(err)
	assert.EqualValues(fiber.StatusUnprocessableEntity, res.StatusCode)
	assert.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
}

func TestErrorHandler_InvalidValidationError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app, _, _ := setupErrorHandlerApp(t)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return &validator.InvalidValidationError{}
	})
	res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.Nil(err)
	assert.EqualValues(fiber.StatusUnprocessableEntity, res.StatusCode)
	assert.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
}

func TestErrorHandler_AnyError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app, _, _ := setupErrorHandlerApp(t)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return errors.New("any other error")
	})
	res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.Nil(err)
	assert.EqualValues(fiber.StatusInternalServerError, res.StatusCode)
	assert.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
}
