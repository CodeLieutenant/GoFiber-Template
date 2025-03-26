package handlers

import (
	"github.com/CodeLieutenant/uberfx-common/v3/http/fiber/fiberfx"

	"github.com/CodeLieutenant/GoFiber-Boilerplate/app/handlers/helloworld"
)

func Handlers() fiberfx.RoutesFx {
	return fiberfx.Routes(
		[]fiberfx.RouteFx{
			fiberfx.Get("/", helloworld.HelloWorld),
		},
		fiberfx.WithPrefix("/"),
	)
}
