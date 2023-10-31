package main

import (
	"github.com/spf13/cobra"

	"github.com/BrosSquad/GoFiber-Boilerplate/app/commands"
	"github.com/BrosSquad/GoFiber-Boilerplate/app/constants"
	"github.com/BrosSquad/GoFiber-Boilerplate/core/cmd"
)

const Version = "0.0.1"

func main() {
	cmd.Execute(Version, []*cobra.Command{
		commands.Serve(),
	}, constants.AppName, constants.AppName, constants.AppDescription)
}
