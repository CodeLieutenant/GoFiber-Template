package config

import (
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/constants"
	"github.com/spf13/viper"

	utilsconfig "github.com/nano-interactive/go-utils/config"
)

type (
	Logging struct {
		Level       string `mapstructure:"level" json:"level" yaml:"level"`
		PrettyPrint bool   `mapstructure:"pretty_print" json:"pretty_print" yaml:"pretty_print"`
	}

	HTTP struct {
		Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
		Port int    `mapstructure:"port" json:"port" yaml:"port"`
	}

	Config struct {
		Logging Logging `mapstructure:"logging" json:"logging" yaml:"logging"`
		HTTP    HTTP    `mapstructure:"http" json:"http" yaml:"http"`
	}
)

func New() (Config, error) {
	cfg := utilsconfig.Config{
		ProjectName: constants.AppName,
		Name:        "config",
		Type:        "yaml",
		Paths: []string{
			"$XDG_CONFIG_HOME/" + constants.AppName,
			"/etc/" + constants.AppName,
			".",
		},
	}

	v, err := utilsconfig.NewWithModifier(cfg)
	if err != nil {
		return Config{}, err
	}

	return NewWithViper(v)
}

func NewWithViper(v *viper.Viper) (Config, error) {
	c := Config{}

	if err := v.Unmarshal(&c); err != nil {
		return Config{}, err
	}

	return c, nil
}
