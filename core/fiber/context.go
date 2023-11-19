package fiber

import (
	"context"

	gofiber "github.com/gofiber/fiber/v2"

	"github.com/BrosSquad/GoFiber-Boilerplate/core/constants"
)

func Context() gofiber.Handler {
	return func(ctx *gofiber.Ctx) error {
		c, cancel := context.WithCancel(context.Background())

		ctx.Locals(constants.CancelFuncContextKey, cancel)
		ctx.SetUserContext(c)

		err := ctx.Next()

		cancelFnWillBeCalled := ctx.Locals(constants.CancelWillBeCalledContextKey)

		if cancelFnWillBeCalled == nil {
			cancel()
		}

		return err
	}
}
