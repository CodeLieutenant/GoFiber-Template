package main

import (
	"github.com/dmalusev/uberfx-common/cmd"

	"github.com/dmalusev/GoFiber-Boilerplate/app/commands"
	"github.com/dmalusev/GoFiber-Boilerplate/app/constants"
)

const Version = "0.0.1"

func main() {
	cmd.Execute(
		Version,
		constants.AppName,
		constants.AppName,
		constants.AppDescription,
		commands.Serve(),
	)
}
