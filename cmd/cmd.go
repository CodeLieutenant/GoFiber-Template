package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/cmd/base"
	"github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/cmd/commands"
)


func registerCommands(root *cobra.Command) {
	root.AddCommand(commands.Serve())
}


func Execute(version string) {
	rootCmd := &cobra.Command{
		Use:                "boilerplate",
		Short:              "boilerplate",
		Long:               `Go Fiber Boilerplate`,
		PersistentPreRunE:  base.LoadConfig,
		PersistentPostRunE: base.CloseResources,
		Version:            version,
	}

	flags := rootCmd.PersistentFlags()

	flags.StringVarP(&base.EnvironmentStr, "env", "e", "production", "Running EnvironmentStr (Production|Development|Testing)")
	flags.StringVarP(&base.ConfigType, "config-type", "t", "yaml", "Configuration Type (yaml|json|toml)")
	flags.StringVarP(&base.ConfigName, "config-name", "c", "config", "Configuration name")
	flags.StringVarP(&base.LoggingLevel, "log-level", "l", "info", "Logging Level (Trace|Debug|Info|Warn|Error|Fatal)")
	flags.BoolVarP(&base.FiberLogo, "fiber-logo", "f", false, "Display Fiber Information")
	flags.BoolVarP(&base.LogPrettyPrint, "log-pretty-print", "p", false, "Pretty print STDOUT/STDERR logs")

	registerCommands(rootCmd)

	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		log.Fatal().
			Err(err).
			Msg("Error while running command")
	}
}
