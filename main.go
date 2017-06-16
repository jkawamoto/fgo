//
// main.go
//
// Copyright (c) 2016-2017 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

package main

import (
	"os"

	"github.com/jkawamoto/fgo/command"
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
	app.Before = command.Prepare
	app.EnableBashCompletion = true
	app.Copyright = `This software is released under the MIT License.
   See https://jkawamoto.github.io/fgo/info/licenses/ for more information.`

	app.Run(os.Args)
}
