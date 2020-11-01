package api

import (
	"context"
	"io"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type (
	Container struct {
		Ctx    context.Context
		DB     *gorm.DB
		Logger zerolog.Logger
	}

	Interface interface {
		io.Closer
		Register(c *Container) error
	}
)
