package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/cobra"

	appLogger "github.com/nano-interactive/go-utils/logging"

	"github.com/BrosSquad/GoFiber-Boilerplate/app/config"
	"github.com/BrosSquad/GoFiber-Boilerplate/app/container"
	"github.com/BrosSquad/GoFiber-Boilerplate/core/constants"
)

func Execute(version string, cmds []*cobra.Command, use, short, long string) {
	rootCmd := &cobra.Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Version: version,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			ctx := context.Background()
			cfg, err := config.New()
			if err != nil {
				return err
			}

			appLogger.ConfigureDefaultLogger(cfg.Logging.Level, cfg.Logging.PrettyPrint)
			//nolint:all
			zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

			cnt := container.New(ctx, cfg)
			ctx = context.WithValue(ctx, constants.ContainerContextKey, cnt)
			ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
			ctx = context.WithValue(ctx, constants.CancelContextKey, cancel)

			cmd.SetContext(ctx)
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cancel := ctx.Value(constants.CancelContextKey).(context.CancelFunc)
			cnt := ctx.Value(constants.ContainerContextKey).(*container.Container)
			cancel()
			return cnt.Close()
		},
	}

	rootCmd.AddCommand(cmds...)

	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		log.Fatal().
			Err(err).
			Msg("Error while running command")
	}
}
