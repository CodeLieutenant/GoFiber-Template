package main

import (
	"github.com/BrosSquad/GoFiber-Boilerplate/app/commands"
	"github.com/BrosSquad/GoFiber-Boilerplate/app/constants"
	"github.com/BrosSquad/GoFiber-Boilerplate/core/cmd"
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
