package container

import (
	"go.uber.org/fx"
)

func New(opts ...fx.Option) *fx.App {
	o := make([]fx.Option, 1, len(opts)+1)
	o[0] = Logger()
	return fx.New(append(o, opts...)...)
}

func Logger() fx.Option {
	return fx.Provide()
}
