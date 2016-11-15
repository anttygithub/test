package main

import (
	"os"

	subscriber "github.com/anttygithub/test/asrc"
	"github.com/urfave/cli"
)

var VERSION = "v0.0.0-dev"

func main() {
	app := cli.NewApp()
	app.Name = "test"
	app.Version = VERSION
	app.Usage = "You need help!"
	app.Action = subscriber.Main
	app.Run(os.Args)
}
