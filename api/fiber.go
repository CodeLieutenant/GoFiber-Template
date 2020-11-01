package api

import (
	"fmt"
	"gohack/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/viper"
)

type fiberApi struct {
	app *fiber.App
}

func NewFiberApi(prefork bool) Interface {
	return fiberApi{
		app: fiber.New(fiber.Config{
			Prefork: prefork,
		}),
	}
}

func (f fiberApi) Register(c *Container) error {
	f.app.Use(logger.New())
	f.app.Use(requestid.New(requestid.Config{
		Generator: utils.UniqueStringGenerator,
	}))

	if err := f.app.Listen(viper.GetString("http.address")); err != nil {
		return fmt.Errorf("Error while starting application: %v", err)
	}
	return nil
}

func (f fiberApi) Close() error {
	return f.app.Shutdown()
}
