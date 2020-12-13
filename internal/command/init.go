package command

import (
	"os"

	"github.com/urfave/cli"
	// "server/internal/logs"
)

var commands []cli.Command
var app *cli.App

func init() {
	app = cli.NewApp()
}

func RunApp() {
	app.Commands = commands
	err := app.Run(os.Args)
	if err != nil {
		// logs.Critical("App run failed: ", err)
	}
}
