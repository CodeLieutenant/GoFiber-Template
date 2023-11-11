package commands

import (
	"github.com/spf13/cobra"

	"github.com/BrosSquad/GoFiber-Boilerplate/app/http"
	"github.com/BrosSquad/GoFiber-Boilerplate/core/container"
	"github.com/BrosSquad/GoFiber-Boilerplate/core/http/httpfx"
)

func Serve() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {
			// ctx := cmd.Context()
			app := container.New(
				httpfx.Module("0.0.0.0:8000", "Test", true, http.Handlers()),
			)
			app.Run()

			// di := ctx.Value(constants.ContainerContextKey).(*container.Container)
			// app := http.CreateApplication(ctx, di, true)

			// cfg := di.GetConfig().HTTP
			// go corehttp.RunServer(cfg.Addr, cfg.Port, app)

			// <-ctx.Done()
			return nil
			// return app.ShutdownWithTimeout(10 * time.Second)
		},
	}
}
