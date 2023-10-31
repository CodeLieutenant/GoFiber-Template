package main

import (
	"github.com/spf13/cobra"

	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/cmd"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/cmd/commands"
)

const Version = "0.0.1"

func main() {
	cmd.Execute(Version, []*cobra.Command{
		commands.Serve(),
	})
}
