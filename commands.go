//
// commands.go
//
// Copyright (c) 2016 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

package main

import (
	"fmt"
	"os"

	"github.com/jkawamoto/fgo/command"
	"github.com/urfave/cli"
)

// GlobalFlags defines global flags.
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "pkg, p",
		Usage: "overwrite directory `NAME` to store package files. (Default: pkg)",
	},
	cli.StringFlag{
		Name:  "brew, b",
		Usage: "overwrite directory `NAME` to store homebrew formula. (Default: homebrew)",
	},
}

// Commands defines sub commands.
var Commands = []cli.Command{
	{
		Name:      "init",
		Usage:     "create Makefile and other related directories.",
		ArgsUsage: "[user name] [repository name]",
		Description: `init command creates a Makefile which will be used to compile your project,
and a template of homebrew formula. If a Makefile or a template of homebrew
formula already exist, fgo won't overwrite them.

To create the template of homebrew formula, a user name and a repository name
in GitHub is required. By default, fgo checks your git configuration to get
those information but you can given them by the arguments.

If your git configuration doesn't have both information and you don't give them
as the arguments, this command will skip to create the template. In this case,
you need to re-run init command after setting git configuration.

You can edit the Makefile and the template of homebrew formula, but build and
release targets are necessary to run build command.`,
		Action: command.CmdInit,
	},
	{
		Name:      "build",
		Usage:     "build binaries, upload them, and update brew formula.",
		ArgsUsage: "[version]",
		Description: `build command runs build and release targets in the Makefile to build your
software and upload the binary files to GitHub. This command takes an argument,
version, which specifies the version to be created. If it is omitted, "snapshot"
will be used and uploading will be skipped.

This command also updates the homebrew formula. After finishing this command,
you need to push the updated formula.`,
		Action: command.CmdBuild,
	},
	{
		Name:      "update",
		Usage:     "update only brew formula.",
		ArgsUsage: "version",
		Description: `update command updates the homebrew formula for a given version. build command
updates the homebrew formula but sometimes you may need to re-update it to a
specific version. This command do that.`,
		Action: command.CmdUpdate,
	},
}

// CommandNotFound defines an action when a given command won't be supported.
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
