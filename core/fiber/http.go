package fiber

import (
	"encoding/json"
	"fmt"
	"net"

	gofiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	fiberutils "github.com/gofiber/fiber/v2/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func RunServer(ip string, port int, app *gofiber.App) {
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

func CreateApplication(appName string, displayInfo bool, logger zerolog.Logger) *gofiber.App {
	var (
		jsonEncoder fiberutils.JSONMarshal   = json.Marshal
		jsonDecoder fiberutils.JSONUnmarshal = json.Unmarshal
	)

	staticConfig := gofiber.Config{
		StrictRouting:                true,
		EnablePrintRoutes:            false,
		Prefork:                      false,
		DisableStartupMessage:        !displayInfo,
		DisableDefaultDate:           true,
		DisableHeaderNormalizing:     false,
		DisablePreParseMultipartForm: true,
		AppName:                      appName,
		ErrorHandler:                 Error(logger),
		JSONEncoder:                  jsonEncoder,
		JSONDecoder:                  jsonDecoder,
	}

	app := gofiber.New(staticConfig)

	app.Use(recover.New())
	app.Use(Context())

	return app
}
