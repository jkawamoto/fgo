//
// command/build.go
//
// Copyright (c) 2016 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

package command

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// BuildOpt defines options for cmdInit.
type BuildOpt struct {
	// Directory to store package files
	Dest string
	// Directory to store brew file
	Brew string
	// Version string.
	Version string
}

func CmdBuild(c *cli.Context) error {

	opt := BuildOpt{
		Dest:    c.String("dest"),
		Brew:    c.String("brew"),
		Version: c.Args().First(),
	}

	// These codes are not necessary but urfave/cli doesn't work.
	if opt.Dest == "" {
		opt.Dest = "pkg"
	}
	if opt.Brew == "" {
		opt.Brew = "brew"
	}

	if err := cmdBuild(&opt); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil

}

func cmdBuild(opt *BuildOpt) (err error) {

	// Build and upload via make.
	fmt.Println(chalk.Bold.TextStyle("Building binaries."))
	if err = build(opt.Version); err != nil {
		return
	}

	return cmdUpdate(opt.Dest, opt.Brew, opt.Version)

}

func build(version string) (err error) {

	var cmd *exec.Cmd
	if version != "" {
		cmd = exec.Command("make", "build", "release", fmt.Sprintf("VERSION=%s", version))
	} else {
		cmd = exec.Command("make", "build")
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	go io.Copy(os.Stdout, stdout)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return
	}
	go io.Copy(os.Stderr, stderr)

	return cmd.Run()

}
