package helloworld

import (
	"github.com/gofiber/fiber/v2"
)

func HelloWorld() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// logger.Info().Msg("Hello World!")
		return c.SendString("Hello, World!")
	}
}
