/*
 * main.go
 *
 * Copyright (c) 2016-2018 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package main

import (
	"os"

	"github.com/mattn/go-colorable"
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
	app.EnableBashCompletion = true
	app.Writer = colorable.NewColorableStdout()
	app.ErrWriter = colorable.NewColorableStderr()
	app.Copyright = `This software is released under the MIT License.
   See https://jkawamoto.github.io/fgo/info/licenses/ for more information.`

	app.Run(os.Args)
}
