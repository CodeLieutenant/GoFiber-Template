package commands

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"github.com/CodeLieutenant/uberfx-common/v3/configfx"
	"github.com/CodeLieutenant/uberfx-common/v3/http/fiber/fiberfx"
	"github.com/CodeLieutenant/uberfx-common/v3/loggerfx"

	"github.com/CodeLieutenant/GoFiber-Boilerplate/app/config"
	"github.com/CodeLieutenant/GoFiber-Boilerplate/app/constants"
	"github.com/CodeLieutenant/GoFiber-Boilerplate/app/handlers"
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
				loggerfx.ZerologModule(loggerSink(&cfg.Logging)),
				fiberfx.App(constants.AppName, handlers.Handlers()),
				fiberfx.RunApp(cfg.HTTP.Addr, constants.AppName, cfg.HTTP.ShutdownTimeout),
			)

			app.Run()
			return nil
		},
	}
}
