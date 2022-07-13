package commands

import (
	"context"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/BrosSquad/GoFiber-Boilerplate/cmd/base"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/constants"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/container"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/http"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/http/handlers"
)

func startHttpServer(ctx context.Context, c *container.Container) {
	app := http.CreateApplication(
		c,
		constants.AppName,
		base.Environment,
		base.FiberLogo,
		base.ViperConfig.GetBool("http.enable_monitor"),
		handlers.Error(
			c.GetLogger(),
			c.GetTranslator(),
		),
	)

	http.RunServer(
		base.ViperConfig.GetString("http.host"),
		base.ViperConfig.GetInt("http.port"),
		app,
	)
}

func Serve() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		RunE: func(command *cobra.Command, args []string) error {
			c := base.GetContainer()

			defer func(c *container.Container) {
				err := c.Close()
				if err != nil {
					log.Error().Err(err).Msg("Failed to close DI Container")
				}
			}(c)

			ctx, cancel := context.WithCancel(command.Context())
			sig := make(chan os.Signal, 1)
			signal.Notify(sig, os.Interrupt)

			defer cancel()

			log.Info().Msg("Starting HTTP Server")
			go startHttpServer(ctx, c)

			<-sig
			cancel()

			return nil
		},
	}
}
