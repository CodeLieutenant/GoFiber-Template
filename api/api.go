package api

import (
	"io"

	"github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/container"
)

type (
	Interface interface {
		io.Closer
		Register(c *container.Container) error
	}
)
