package container

import (
	"github.com/rs/zerolog"

	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/logging"
)

func (c *Container) GetLogger() zerolog.Logger {
	return logging.New(c.loggingLevel, c.loggingPrettyPrint)
}
