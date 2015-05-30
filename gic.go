package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gic"
	app.Version = "0.0.1"
	app.Usage = "Create GitHub issue from template."
	app.Author = "i2bskn"
	app.Email = "i2bskn@gmail.com"
	app.Commands = Commands
	app.Run(os.Args)
}
