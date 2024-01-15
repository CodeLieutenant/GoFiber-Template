package handlers

import (
	"github.com/dmalusev/GoFiber-Boilerplate/app/handlers/helloworld"
	"github.com/dmalusev/uberfx-common/fiber/fiberfx"
)

func Handlers() fiberfx.RoutesFx {
	return fiberfx.Routes(
		fiberfx.WithRoutes(fiberfx.Get("/", helloworld.HelloWorld)),
		fiberfx.WithPrefix("/"),
	)
}
