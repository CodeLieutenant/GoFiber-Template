package commands

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"github.com/BrosSquad/GoFiber-Boilerplate/app/config"
	"github.com/BrosSquad/GoFiber-Boilerplate/app/constants"
	"github.com/BrosSquad/GoFiber-Boilerplate/app/handlers"
	"github.com/BrosSquad/GoFiber-Boilerplate/core/configfx"
	"github.com/BrosSquad/GoFiber-Boilerplate/core/fiber/fiberfx"
	"github.com/BrosSquad/GoFiber-Boilerplate/core/loggerfx"
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
