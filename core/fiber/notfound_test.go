package fiber_test

// func setupNotFoundApplication() *fiber.App {
// 	app := fiber.New()

// 	app.Use(handlers.NotFound())

// 	return app
// }

// func TestNotFound_JsonResponse(t *testing.T) {
// 	t.Parallel()
// 	assert := require.New(t)

// 	app := setupNotFoundApplication()

// 	res := testing_utils.Get(app, "/", testing_utils.WithHeaders(http.Header{
// 		fiber.HeaderAccept: []string{fiber.MIMEApplicationJSONCharsetUTF8},
// 	}))

// 	assert.Equal(http.StatusNotFound, res.StatusCode)
// 	assert.Equal(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))

// 	var errRes handlers.ErrorResponse
// 	_ = json.NewDecoder(res.Body).Decode(&errRes)

// 	assert.EqualValues(handlers.ErrorResponse{
// 		Message: "Page is not found",
// 	}, errRes)
// }
