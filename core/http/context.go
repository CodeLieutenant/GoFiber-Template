package http

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/BrosSquad/GoFiber-Boilerplate/core/constants"
)

func Context(base context.Context) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		c, cancel := context.WithCancel(base)

		ctx.Locals(constants.CancelFuncContextKey, cancel)
		ctx.SetUserContext(c)

		err := ctx.Next()

		cancelFnWillBeCalled := ctx.Locals(constants.CancelWillBeCalledContextKey)

		if cancelFnWillBeCalled == nil {
			defer cancel()
		}

		return err
	}
}
