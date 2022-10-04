package commands

import (
	"context"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/nano-interactive/GoFiber-Boilerplate/cmd/base"
	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/constants"
	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/container"
	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/http"
	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/http/handlers"
)

func startHttpServer(ctx context.Context, c *container.Container) {
	app := http.CreateApplication(
		c,
		constants.AppName,
		base.Environment,
		base.FiberLogo,
		c.GetConfig().HTTP.EnableMonitor,
		handlers.Error(c.GetLogger(), c.GetTranslator()),
	)

	http.RunServer(
		c.GetConfig().HTTP.Host,
		c.GetConfig().HTTP.Port,
		app,
	)
}

func Serve() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		RunE: func(command *cobra.Command, args []string) error {
			return execute(command)
		},
	}
}

func execute(command *cobra.Command) error {
	ctx, cancel := context.WithCancel(command.Context())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	defer cancel()

	c := base.GetContainer(ctx)

	defer func(c *container.Container) {
		err := c.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close DI Container")
		}
	}(c)

	log.Info().Msg("Starting HTTP Server")
	go startHttpServer(ctx, c)

	<-sig
	cancel()

	return nil
}
