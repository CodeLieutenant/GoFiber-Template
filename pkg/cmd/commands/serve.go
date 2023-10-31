package commands

import (
	"context"
	"time"

	"github.com/spf13/cobra"

	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/constants"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/container"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/http"
)

func Serve() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			di := ctx.Value(constants.ContainerContextKey).(*container.Container)
			app := http.CreateApplication(ctx, di, true)
			cfg := di.GetConfig().HTTP

			go http.RunServer(cfg.Addr, cfg.Port, app)

			<-ctx.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			return app.ShutdownWithContext(ctx)
		},
	}
}
