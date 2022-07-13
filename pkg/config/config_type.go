package config

import (
	"errors"
	"strings"
)

type ConfigType string

const (
	JSON ConfigType = "json"
	YAML ConfigType = "yaml"
	TOML ConfigType = "toml"
)

func ParseConfigType(configType string) (ConfigType, error) {
	switch strings.ToLower(configType) {
	case "json":
		return JSON, nil
	case "yaml", "": // Empty string as default
		return YAML, nil
	case "toml":
		return TOML, nil
	default:
		return "", errors.New("Invalid Configuration Type: JSON, YAML, TOML or \"\"(empty string), Given: " + configType)
	}
}
