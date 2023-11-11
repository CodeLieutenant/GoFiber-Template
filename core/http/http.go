package http

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/BrosSquad/GoFiber-Boilerplate/app/container"
	"github.com/BrosSquad/GoFiber-Boilerplate/core/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	fiberutils "github.com/gofiber/fiber/v2/utils"
	"github.com/nano-interactive/go-utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Routes func(fiber.Router, *container.Container)

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

func CreateApplication(appName string, displayInfo bool) *fiber.App {
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
		AppName:                      appName,
		ErrorHandler:                 Error(zerolog.Nop()),
		JSONEncoder:                  jsonEncoder,
		JSONDecoder:                  jsonDecoder,
	}

	app := fiber.New(staticConfig)

	app.Use(recover.New())
	app.Use(Context())
	app.Use(requestid.New(requestid.Config{
		Generator: func() string {
			return utils.RandomString(32)
		},
		ContextKey: constants.RequestIDContextKey,
	}))

	return app
}
