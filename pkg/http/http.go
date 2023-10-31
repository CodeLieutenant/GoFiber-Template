package http

import (
	"context"
	"fmt"
	"net"

	"github.com/goccy/go-json"
	"github.com/nano-interactive/go-utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/constants"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/container"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/http/handlers"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/http/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	fiberutils "github.com/gofiber/fiber/v2/utils"
)

func CreateApplication(ctx context.Context, c *container.Container, displayInfo bool) *fiber.App {
	var (
		jsonEncoder fiberutils.JSONMarshal   = json.Marshal
		jsonDecoder fiberutils.JSONUnmarshal = json.Unmarshal
	)

	staticConfig := fiber.Config{
		StrictRouting:                true,
		EnablePrintRoutes:            false,
		Prefork:                      false,
		DisableStartupMessage:        !displayInfo,
		DisableDefaultDate:           true,
		DisableHeaderNormalizing:     false,
		DisablePreParseMultipartForm: true,
		AppName:                      constants.AppName,
		ErrorHandler:                 handlers.Error(zerolog.Nop()),
		JSONEncoder:                  jsonEncoder,
		JSONDecoder:                  jsonDecoder,
	}

	app := fiber.New(staticConfig)

	app.Use(recover.New())
	app.Use(middleware.Context(ctx))
	app.Use(requestid.New(requestid.Config{
		Generator: func() string {
			return utils.RandomString(32)
		},
		ContextKey: constants.RequestIDContextKey,
	}))

	registerHandlers(app, c)

	return app
}

func RunServer(ip string, port int, app *fiber.App) {
	addr := fmt.Sprintf("%s:%d", ip, port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.
			Fatal().
			Err(err).
			Msg("Error while creating net.Listener for HTTP Server")
	}

	err = app.Listener(listener)

	if err != nil {
		log.
			Fatal().
			Err(err).
			Msg("Cannot start Fiber HTTP Server")
	}
}
