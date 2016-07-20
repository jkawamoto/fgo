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
		Name:  "dest, d",
		Value: "pkg",
		Usage: "directory `NAME` to store package files",
	},
	cli.StringFlag{
		Name:  "brew, b",
		Value: "brew",
		Usage: "directory `NAME` to store brew file",
	},
}

// TODO: Main command takes version name.

// Commands defines sub commands.
var Commands = []cli.Command{
	{
		Name:      "init",
		Usage:     "create Makefile and other related directories.",
		ArgsUsage: "[username]",
		Action:    command.CmdInit,
		Flags:     []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
