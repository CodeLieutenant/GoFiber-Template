package config

import (
	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/constants"
	utilsConfig "github.com/nano-interactive/go-utils/config"
	"github.com/spf13/viper"
)

type Config struct {
	HTTP struct {
		Host          string `mapstructure:"host"`
		Port          int    `mapstructure:"port"`
		Domain        string `mapstructure:"domain"`
		EnableMonitor bool   `mapstructure:"enable_monitor"`
	} `mapstructure:"http"`

	CORS struct {
		Headers []string `mapstructure:"headers"`
		Origins []string `mapstructure:"origins"`
		Methods []string `mapstructure:"methods"`
	} `mapstructure:"cors"`

	Database struct {
		Redis struct {
			Host            string `mapstructure:"host"`
			Port            int    `mapstructure:"port"`
			Password        string `mapstructure:"password"`
			DefaultDatabase int    `mapstructure:"database"`
			CSRF            struct {
				Database int `mapstructure:"db"`
			} `mapstructure:"csrf"`
		} `mapstructure:"redis"`
	} `mapstructure:"database"`
}

func NewWithEnvironment(environment, configName, configType string) (*Config, error) {
	cfg := utilsConfig.Config{
		ProjectName: constants.AppName,
		Env:         environment,
		Name:        configName,
		Type:        configType,
	}

	v, err := utilsConfig.NewWithModifier(cfg, ConfigureEnv, EnvironmentalVariables)

	if err != nil {
		return nil, err
	}

	return NewWithViper(v)
}

func NewWithViper(v *viper.Viper) (*Config, error) {
	c := new(Config)

	if err := v.Unmarshal(c); err != nil {
		return nil, err
	}

	return c, nil
}
