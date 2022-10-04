package hello_world_test

import (
	"net/http"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	nanoTesting "github.com/nano-interactive/go-utils/testing"
	nanoTestingFiber "github.com/nano-interactive/go-utils/testing/fiber"
	nanoTestingHttp "github.com/nano-interactive/go-utils/testing/http"

	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/http/handlers/hello_world"
	"github.com/nano-interactive/GoFiber-Boilerplate/testing_utils"
)

func TestHelloWorldHandler(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app, _ := testing_utils.CreateApplication(t)
	_, loggerAssert :=  nanoTesting.NewAppTestLogger(t, zerolog.InfoLevel)

	client := nanoTestingFiber.New[any](t, app, false)

	app.Get("/", hello_world.HelloWorld(loggerAssert.Logger()))

	res := nanoTestingHttp.Get(t, client, "/")

	assert.Equal(http.StatusOK, res.StatusCode)
}
