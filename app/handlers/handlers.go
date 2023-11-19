package handlers

import (
	"github.com/BrosSquad/GoFiber-Boilerplate/app/handlers/helloworld"
	"github.com/BrosSquad/GoFiber-Boilerplate/core/fiber/fiberfx"
)

func Handlers() fiberfx.RoutesFx {
	return fiberfx.Routes(
		"/",
		nil,
		fiberfx.Get("/", helloworld.HelloWorld),
	)
}
