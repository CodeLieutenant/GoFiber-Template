package middleware

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"

	"github.com/stretchr/testify/require"

	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/constants"
)

func TestContextMiddleware(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app := fiber.New()

	app.Use(Context)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	h := app.Handler()

	ctx := &fasthttp.RequestCtx{}

	h(ctx)

	assert.NotNil(ctx.UserValue(constants.CancelFuncContextKey))
}
