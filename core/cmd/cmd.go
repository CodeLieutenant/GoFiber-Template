package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func Execute(version string, use, short, long string, cmds ...*cobra.Command) {
	rootCmd := &cobra.Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Version: version,
	}

	rootCmd.AddCommand(cmds...)

	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		log.Fatal().
			Err(err).
			Msg("Error while running command")
	}
}
