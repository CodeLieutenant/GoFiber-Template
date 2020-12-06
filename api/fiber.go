package api

import (
	"github.com/pkg/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"

	"github.com/spf13/viper"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type fiberAPI struct {
	app   *fiber.App
	debug bool
}

func NewFiberAPI(prefork, debug bool, errorHandler fiber.ErrorHandler) Interface {
	return fiberAPI{
		app: fiber.New(fiber.Config{
			Prefork:      prefork,
			ErrorHandler: errorHandler,
		}),
		debug: debug,
	}
}

func (f fiberAPI) Register(_ *Container) error {
	f.app.Use(logger.New())

	if !f.debug {
		f.app.Use(recover.New())
		f.app.Use(compress.New(compress.Config{
			Next:  nil,
			Level: compress.LevelBestSpeed,
		}))
	} else {
		f.app.Use(pprof.New())
		f.app.Get("/monitor", monitor.New())
	}

	if err := f.app.Listen(viper.GetString("http.address")); err != nil {
		return errors.Wrap(err, "Error while starting application")
	}

	return nil
}

func (f fiberAPI) Close() error {
	return f.app.Shutdown()
}
