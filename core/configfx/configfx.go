package configfx

import (
	utilsconfig "github.com/nano-interactive/go-utils/config"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func New[T any](appName string) (T, error) {
	v, err := utilsconfig.NewWithModifier(utilsconfig.Config{
		ProjectName: appName,
		Name:        "config",
		Type:        "yaml",
		Paths: []string{
			"$XDG_CONFIG_HOME/" + appName,
			"/etc/" + appName,
			".",
		},
	})
	if err != nil {
		var t T
		return t, err
	}

	return NewWithViper[T](v)
}

func NewWithViper[T any](v *viper.Viper) (T, error) {
	var c T

	if err := v.Unmarshal(&c); err != nil {
		return c, err
	}

	return c, nil
}

func Module[T any](cfg T) fx.Option {
	return fx.Module(
		"config",
		fx.Supply(cfg),
	)
}
