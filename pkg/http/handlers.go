package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/config"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/container"

	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/http/handlers/hello_world"
)

func registerHandlers(app *fiber.App, c *container.Container, environment config.Env) {
	app.Get("/", hello_world.HelloWorld(c.GetLogger()))
}
