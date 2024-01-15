package commands

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"github.com/dmalusev/uberfx-common/configfx"
	"github.com/dmalusev/uberfx-common/fiber/fiberfx"
	"github.com/dmalusev/uberfx-common/loggerfx"

	"github.com/dmalusev/GoFiber-Boilerplate/app/config"
	"github.com/dmalusev/GoFiber-Boilerplate/app/constants"
	"github.com/dmalusev/GoFiber-Boilerplate/app/handlers"
)

func loggerSink(cfg *config.Logging) loggerfx.Sink {
	return loggerfx.Sink{
		Level:       cfg.Level,
		Type:        loggerfx.Stdout,
		PrettyPrint: cfg.PrettyPrint,
	}
}

func Serve() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		RunE: func(_ *cobra.Command, _ []string) error {
			cfg, err := configfx.New[config.Config](constants.AppName)
			if err != nil {
				return err
			}

			app := fx.New(
				configfx.Module(cfg),
				loggerfx.Module(loggerSink(&cfg.Logging)),
				fiberfx.App(constants.AppName, cfg.App.FiberInfo, handlers.Handlers()),
				fiberfx.RunApp(cfg.HTTP.Addr, constants.AppName, cfg.HTTP.ShutdownTimeout),
			)

			app.Run()
			return nil
		},
	}
}
