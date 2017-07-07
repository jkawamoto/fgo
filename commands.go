//
// commands.go
//
// Copyright (c) 2016-2017 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/jkawamoto/fgo/command"
	"github.com/urfave/cli"
)

const (
	// DefaultPackageDir defines default package directory.
	DefaultPackageDir = "pkg"
	// DefaultHomebrewDir defines default homebrew formula directory.
	DefaultHomebrewDir = "homebrew"
)

// GlobalFlags defines global flags.
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "pkg, p",
		Usage: "directory `NAME` to store package files",
		Value: DefaultPackageDir,
	},
	cli.StringFlag{
		Name:  "brew, b",
		Usage: "directory `NAME` to store homebrew formula",
		Value: DefaultHomebrewDir,
	},
}

// Commands defines sub commands.
var Commands = []cli.Command{
	{
		Name:      "init",
		Usage:     "create Makefile and other related directories.",
		ArgsUsage: "[user name [repository name]]",
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
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "desc, d",
				Usage: "Set `TEXT` to description section of the formula template",
			},
		},
	},
	{
		Name:      "build",
		Usage:     "build binaries, upload them, and update the brew formula.",
		ArgsUsage: "[version]",
		Description: `build command runs build and release targets in the Makefile to build your
software and upload the binary files to GitHub. This command takes an argument,
version, which specifies the version to be created. If it is omitted, "snapshot"
will be used and uploading will be skipped.

To run this command, a GitHub API token is required. Users have to give a token
via one of the -t/--token flag, GITHUB_TOKEN environment variable, and github.token
variable in your .gitconfig.

If -b/--body flag isn't given but your CHANGELOG.md contains a release note
associated with that version, the release note will be copied to the release
page in GitHub.

This command also updates the homebrew formula. After finishing this command,
you need to push the updated formula.`,
		Action: command.CmdBuild,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "t, token",
				Usage:  "GitHub API `TOKEN` for uploading binaries",
				EnvVar: "GITHUB_TOKEN",
			},
			cli.StringFlag{
				Name:  "b, body",
				Usage: "`TEXT` describing the contents of this release",
			},
			cli.IntFlag{
				Name:  "p, process",
				Usage: "the number of goroutines",
				Value: runtime.NumCPU(),
			},
			cli.BoolFlag{
				Name:  "delete",
				Usage: "delete release and its git tag in advance if exists",
			},
			cli.BoolFlag{
				Name:  "draft",
				Usage: "create a draft (unpublished) release",
			},
			cli.BoolFlag{
				Name:  "pre",
				Usage: "mark this release is a prerelease",
			},
		},
	},
	{
		Name:      "update",
		Usage:     "update the brew formula.",
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
