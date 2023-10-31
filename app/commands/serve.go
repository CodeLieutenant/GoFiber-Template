package commands

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/BrosSquad/GoFiber-Boilerplate/app/container"
	"github.com/BrosSquad/GoFiber-Boilerplate/app/http"
	"github.com/BrosSquad/GoFiber-Boilerplate/core/constants"
	corehttp "github.com/BrosSquad/GoFiber-Boilerplate/core/http"
)

func Serve() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			di := ctx.Value(constants.ContainerContextKey).(*container.Container)
			app := http.CreateApplication(ctx, di, true)

			cfg := di.GetConfig().HTTP
			go corehttp.RunServer(cfg.Addr, cfg.Port, app)

			<-ctx.Done()
			return app.ShutdownWithTimeout(10 * time.Second)
		},
	}
}
