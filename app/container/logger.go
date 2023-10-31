package container

import (
	"io"
	"os"

	appLogger "github.com/nano-interactive/go-utils/logging"
	"github.com/rs/zerolog"
)

func (c *Container) GetLogger() zerolog.Logger {
	var stdout io.Writer = os.Stdout

	if c.config.Logging.PrettyPrint {
		stdout = zerolog.NewConsoleWriter()
	}

	writer := stdout

	return appLogger.New(c.config.Logging.Level, c.config.Logging.PrettyPrint).
		Output(writer).
		With().
		Stack().
		Logger()
}
