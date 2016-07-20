package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "Junpei Kawamoto"
	app.Email = "kawamoto.junpei@gmail.com"
	app.Usage = "Build, upload, and create brew formula for golang application."

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
