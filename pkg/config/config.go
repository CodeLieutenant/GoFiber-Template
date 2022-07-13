package config

import (
	"github.com/spf13/viper"
)

func New(envStr, configName, configTypeStr string) (*viper.Viper, error) {
	env, err := ParseEnvironment(envStr)

	if err != nil {
		return nil, err
	}

	configType, err := ParseConfigType(configTypeStr)

	if err != nil {
		return nil, err
	}

	v := viper.New()

	v.SetConfigName(configName)
	v.SetConfigType(string(configType))

	if env == Production {
		v.AddConfigPath("/etc/sitemap")
		v.AddConfigPath(".")
	} else {
		v.AddConfigPath(".")
	}

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}
