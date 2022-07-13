package hello_world_test

import (
	"net/http"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/http/handlers/hello_world"
	"github.com/BrosSquad/GoFiber-Boilerplate/testing_utils"
)

func TestHelloWorldHandler(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app, _ := testing_utils.CreateApplication()
	_, loggerAssert := testing_utils.NewTestLogger(t, zerolog.InfoLevel)

	app.Get("/", hello_world.HelloWorld(loggerAssert.Logger()))

	res := testing_utils.Get(app, "/")

	assert.Equal(http.StatusOK, res.StatusCode)
}
