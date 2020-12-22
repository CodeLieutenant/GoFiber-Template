package container

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

// TODO: Rename
type Container struct {
	Ctx       context.Context
	Logger    *zerolog.Logger
	DB        *gorm.DB
	Validator *validator.Validate
}
