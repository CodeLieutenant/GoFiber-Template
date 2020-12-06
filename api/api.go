package api

import (
	"context"
	"github.com/go-playground/validator/v10"
	"io"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type (
	Container struct {
		Ctx    context.Context
		DB     *gorm.DB
		Logger zerolog.Logger
		Validator *validator.Validate
	}

	Interface interface {
		io.Closer
		Register(c *Container) error
	}
)
