package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/container"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/http/handlers/helloworld"
)

func registerHandlers(app *fiber.App, c *container.Container) {
	app.Get("/", helloworld.HelloWorld(c.GetLogger()))
}
