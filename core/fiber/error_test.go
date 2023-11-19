package fiber_test

// func setupErrorHandlerApp(t *testing.T) (*fiber.App, *zltest.Tester) {
// 	// v, translations := testing_utils.GetValidator()

// 	logger, loggerTest := nanotesting.NewAppTestLogger(t, zerolog.InfoLevel)

// 	app := fiber.New(fiber.Config{
// 		ErrorHandler: handlers.Error(logger, nil),
// 	})

// 	return app, v, loggerTest
// }

// func TestErrorHandler_ReturnFiberError(t *testing.T) {
// 	t.Parallel()
// 	assert := require.New(t)

// 	app, _, _ := setupErrorHandlerApp(t)

// 	app.Get("/", func(ctx *fiber.Ctx) error {
// 		return fiber.ErrBadGateway
// 	})

// 	m := struct {
// 		Message string `json:"message"`
// 	}{}
// 	res := testing_utils.Get(app, "/")

// 	assert.EqualValues(fiber.StatusBadGateway, res.StatusCode)
// 	assert.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
// 	assert.Nil(json.NewDecoder(res.Body).Decode(&m))
// 	assert.NotEmpty(m.Message)
// }

// func TestErrorHandler_InvalidPayloadError(t *testing.T) {
// 	t.Parallel()
// 	assert := require.New(t)

// 	app, _, _ := setupErrorHandlerApp(t)
// 	app.Get("/", func(ctx *fiber.Ctx) error {
// 		return handlers.ErrInvalidPayload
// 	})

// 	res := testing_utils.Get(app, "/")

// 	m := struct {
// 		Message string `json:"message"`
// 	}{}

// 	assert.EqualValues(fiber.StatusBadRequest, res.StatusCode)
// 	assert.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
// 	assert.Nil(json.NewDecoder(res.Body).Decode(&m))
// 	assert.NotEmpty(m.Message)
// 	assert.Equal(handlers.ErrInvalidPayload.Error(), m.Message)
// }

// func TestErrorHandler_ValidationError(t *testing.T) {
// 	t.Parallel()
// 	assert := require.New(t)
// 	app, _, _ := setupErrorHandlerApp(t)
// 	app.Get("/", func(ctx *fiber.Ctx) error {
// 		return validator.ValidationErrors{}
// 	})
// 	res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
// 	assert.Nil(err)
// 	assert.EqualValues(fiber.StatusUnprocessableEntity, res.StatusCode)
// 	assert.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
// }

// func TestErrorHandler_InvalidValidationError(t *testing.T) {
// 	t.Parallel()
// 	assert := require.New(t)

// 	app, _, _ := setupErrorHandlerApp(t)
// 	app.Get("/", func(ctx *fiber.Ctx) error {
// 		return &validator.InvalidValidationError{}
// 	})
// 	res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
// 	assert.Nil(err)
// 	assert.EqualValues(fiber.StatusUnprocessableEntity, res.StatusCode)
// 	assert.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
// }

// func TestErrorHandler_AnyError(t *testing.T) {
// 	t.Parallel()
// 	assert := require.New(t)
// 	app, _, _ := setupErrorHandlerApp(t)
// 	app.Get("/", func(ctx *fiber.Ctx) error {
// 		return errors.New("any other error")
// 	})
// 	res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
// 	assert.Nil(err)
// 	assert.EqualValues(fiber.StatusInternalServerError, res.StatusCode)
// 	assert.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
// }
