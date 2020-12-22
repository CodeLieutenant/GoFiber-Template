package api

import (
	"github.com/pkg/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/container"
)

type RegisterRoutesHandler func(*container.Container, *fiber.App)

type fiberAPI struct {
	app            *fiber.App
	address        string
	registerRoutes RegisterRoutesHandler
	debug          bool
}

func NewFiberAPI(address string, prefork, debug bool, errorHandler fiber.ErrorHandler, register RegisterRoutesHandler) Interface {
	return fiberAPI{
		app: fiber.New(fiber.Config{
			Prefork:      prefork,
			ErrorHandler: errorHandler,
		}),
		debug:          debug,
		address:        address,
		registerRoutes: register,
	}
}

func (f fiberAPI) Register(c *container.Container) error {

	if !f.debug {
		c.Logger.Debug().Msg("Running in production mode, recover and compression middleware are enabled")
		f.app.Use(recover.New())
		f.app.Use(compress.New(compress.Config{
			Next:  nil,
			Level: compress.LevelBestSpeed,
		}))
	} else {
		c.Logger.Debug().Msg("Running in DEBUG mode, PProf and Monitor (GET /monitor) are enabled")
		f.app.Use(pprof.New())
		f.app.Get("/monitor", monitor.New())
		f.app.Use(logger.New())
	}

	if fiber.IsChild() {
		c.Logger.Debug().Msg("Starting the preforked process")
	} else {
		c.Logger.Debug().Msg("Starting the main application")
	}

	if f.registerRoutes != nil {
		c.Logger.Debug().Msg("Loading the routes")
		f.registerRoutes(c, f.app)
	}

	c.Logger.Debug().Msgf("Listening on address %s", f.address)
	if err := f.app.Listen(f.address); err != nil {
		return errors.Wrap(err, "Error while starting application")
	}

	return nil
}

func (f fiberAPI) Close() error {
	return f.app.Shutdown()
}
