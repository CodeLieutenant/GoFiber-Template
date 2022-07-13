package container

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"

	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/config"
)

type Container struct {
	config       *viper.Viper
	loggingLevel string
	loggingPrettyPrint bool
	environment  config.Env

	validator  *validator.Validate
	translator ut.Translator
}

func New(config *viper.Viper, loggingPrettyPrint bool, loggingLevel string, env config.Env) *Container {
	return &Container{
		config:       config,
		loggingLevel: loggingLevel,
		loggingPrettyPrint: loggingPrettyPrint,
		environment:  env,
	}
}

func (c *Container) GetEnvironment() config.Env {
	return c.environment
}

func (c *Container) Close() error {
	return nil
}
