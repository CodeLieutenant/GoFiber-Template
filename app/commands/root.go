package commands

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/CodeLieutenant/GoFiber-Boilerplate/app/constants"
)

func Execute(version string) error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)

	rootCmd := &cobra.Command{
		Use:     constants.AppName,
		Short:   constants.AppName,
		Long:    constants.AppDescription,
		Version: version,
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			cmd.SetVersionTemplate("Version: " + version)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			cancel()
		},
	}

	rootCmd.AddCommand(Serve())

	return rootCmd.ExecuteContext(ctx)
}
