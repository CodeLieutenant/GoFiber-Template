package container

import (
	"github.com/rs/zerolog"

	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/logging"
)

func (c *Container) GetLogger() zerolog.Logger {
	return logging.New(c.loggingLevel, c.loggingPrettyPrint)
}
