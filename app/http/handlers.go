package http

import (
	"go.uber.org/fx"

	"github.com/BrosSquad/GoFiber-Boilerplate/app/http/helloworld"
	"github.com/BrosSquad/GoFiber-Boilerplate/core/http/httpfx"
)

func Handlers() fx.Option {
	return httpfx.Routes(
		"/",
		httpfx.Get("/", helloworld.HelloWorld),
	)
}
