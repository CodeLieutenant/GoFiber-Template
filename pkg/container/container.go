package container

import (
	"context"

	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/config"
)

type Container struct {
	config *config.Config
	ctx    context.Context
}

func New(ctx context.Context, config config.Config) *Container {
	return &Container{
		ctx:    ctx,
		config: &config,
	}
}

func (c *Container) GetConfig() config.Config {
	return *c.config
}

func (c *Container) Close() error {
	return nil
}
