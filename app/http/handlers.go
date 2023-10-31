package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/BrosSquad/GoFiber-Boilerplate/app/container"
	"github.com/BrosSquad/GoFiber-Boilerplate/app/http/helloworld"
)

func routes(app fiber.Router, c *container.Container) {
	app.Get("/", helloworld.HelloWorld(c.GetLogger()))
}
