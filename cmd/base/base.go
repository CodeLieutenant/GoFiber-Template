package base

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/config"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/container"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/logging"
)

var (
	ViperConfig *viper.Viper
	Environment config.Env

	EnvironmentStr string
	ConfigName string
	ConfigType string
	LoggingLevel   string

	FiberLogo bool
	LogPrettyPrint bool
)

func LoadConfig(*cobra.Command, []string) error {
	logging.ConfigureDefaultLogger(LoggingLevel, LogPrettyPrint)

	v, err := config.New(EnvironmentStr, ConfigName, ConfigType)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load Configuration")
		return err
	}

	ViperConfig = v

	return nil
}

func GetContainer() *container.Container {
	return container.New(ViperConfig,LogPrettyPrint, LoggingLevel, Environment)
}

func CloseResources(*cobra.Command, []string) error {
	return nil
}
