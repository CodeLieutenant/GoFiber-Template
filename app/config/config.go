package config

type (
	App struct {
		FiberInfo bool `mapstructure:"fiber_info" json:"fiber_info" yaml:"fiber_info"`
	}
	Logging struct {
		Level       string `mapstructure:"level" json:"level" yaml:"level"`
		PrettyPrint bool   `mapstructure:"pretty_print" json:"pretty_print" yaml:"pretty_print"`
	}

	Config struct {
		Logging Logging `mapstructure:"logging" json:"logging" yaml:"logging"`
		HTTP    HTTP    `mapstructure:"http" json:"http" yaml:"http"`
		App     App     `mapstructure:"app" json:"app" yaml:"app"`
	}
)
