package config

import (
	"time"
)

type HTTP struct {
	Addr            string        `mapstructure:"addr" json:"addr" yaml:"addr"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout" json:"shutdown_timeout" yaml:"shutdown_timeout"`
}
