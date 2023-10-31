package helloworld_test

// func TestHelloWorldHandler(t *testing.T) {
// 	t.Parallel()
// 	assert := require.New(t)

// 	app, _ := testing_utils.CreateApplication()
// 	_, loggerAssert := testing_utils.NewTestLogger(t, zerolog.InfoLevel)

// 	app.Get("/", hello_world.HelloWorld(loggerAssert.Logger()))

// 	res := testing_utils.Get(app, "/")

// 	assert.Equal(http.StatusOK, res.StatusCode)
// }
