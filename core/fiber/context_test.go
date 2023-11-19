package fiber_test

import (
	"testing"

	gofiber "github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"

	"github.com/stretchr/testify/require"

	"github.com/BrosSquad/GoFiber-Boilerplate/core/constants"
	"github.com/BrosSquad/GoFiber-Boilerplate/core/fiber"
)

func TestContextMiddleware(t *testing.T) {
	t.Parallel()

	assert := require.New(t)
	app := gofiber.New()
	app.Use(fiber.Context())
	app.Get("/", func(ctx *gofiber.Ctx) error {
		return ctx.SendStatus(gofiber.StatusOK)
	})

	h := app.Handler()
	ctx := &fasthttp.RequestCtx{}
	h(ctx)

	assert.NotNil(ctx.UserValue(constants.CancelFuncContextKey))
}
