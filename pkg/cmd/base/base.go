package base

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/nano-interactive/go-utils/environment"
	"github.com/nano-interactive/go-utils/logging"

	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/config"
	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/container"
)

var (
	Environment environment.Env

	EnvironmentStr string
	ConfigName     string
	ConfigType     string
	LoggingLevel   string

	LogPrettyPrint bool
)

func LoadConfig(*cobra.Command, []string) (err error) {
	logging.ConfigureDefaultLogger(LoggingLevel, LogPrettyPrint)

	Environment, err = environment.Parse(EnvironmentStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse environment")
		return err
	}

	return nil
}


func GetContainer(ctx context.Context) *container.Container {
	cfg, err := config.NewWithEnvironment(EnvironmentStr, ConfigName, ConfigType)
	if err != nil {
		log.Fatal().Msg("Failed to load configuration." + err.Error())
	}

	return container.New(ctx, cfg, LogPrettyPrint, LoggingLevel, Environment)
}

func CloseResources(*cobra.Command, []string) error {
	return nil
}
