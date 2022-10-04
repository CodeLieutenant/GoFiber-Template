package container

import (
	"context"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	"github.com/nano-interactive/go-utils/environment"

	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/config"
)

type Container struct {
	ctx context.Context

	config      *config.Config
	environment environment.Env

	loggingLevel       string
	loggingPrettyPrint bool

	validator  *validator.Validate
	translator ut.Translator
}

func New(ctx context.Context, config *config.Config, loggingPrettyPrint bool, loggingLevel string, env environment.Env) *Container {
	return &Container{
		config:             config,
		loggingLevel:       loggingLevel,
		loggingPrettyPrint: loggingPrettyPrint,
		environment:        env,
	}
}

func (c *Container) GetEnvironment() environment.Env {
	return c.environment
}

func (c *Container) GetConfig() *config.Config {
	return c.config
}

func (c *Container) GetBaseContext() context.Context {
	return c.ctx
}

func (c *Container) Close() error {
	return nil
}
