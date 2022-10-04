package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/nano-interactive/go-utils/environment"

	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/container"
	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/http/handlers/hello_world"
)

func registerHandlers(app *fiber.App, c *container.Container, environment environment.Env) {
	app.Get("/", hello_world.HelloWorld(c.GetLogger()))
}
