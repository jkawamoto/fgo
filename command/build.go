/*
 * build.go
 *
 * Copyright (c) 2016-2018 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package command

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/jkawamoto/fgo/fgo"
	colorable "github.com/mattn/go-colorable"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// BuildOpt defines options for cmdInit.
type BuildOpt struct {
	// Directory configurations.
	Directories
	// Version string.
	Version string
	// Options for ghr command.
	GHROpt ghrOpt
}

// ghrOpt defines options for ghr command.
type ghrOpt struct {
	// GitHub API token.
	Token string
	// New release's message body.
	Body string
	// The number of goroutines used in ghr.
	Process int
	// If true and the given version is already released, delete it and create a
	// new release.
	Delete bool
	// If true, the new release won't be published.
	Draft bool
	// If true, the new release will be marked as a prerelease.
	Pre bool
}

// String creates a string representing flags of this options.
func (o *ghrOpt) String() string {

	var opts []string
	if o.Token != "" {
		opts = append(opts, fmt.Sprint("-t ", o.Token))
	}
	if o.Body != "" {
		opts = append(opts, fmt.Sprintf(`-b "%v"`, o.Body))
	}
	if o.Process > 0 {
		opts = append(opts, fmt.Sprint("-p ", o.Process))
	}
	if o.Delete {
		opts = append(opts, "-delete")
	}
	if o.Draft {
		opts = append(opts, "-draft")
	}
	if o.Pre {
		opts = append(opts, "-prerelease")
	}

	return strings.Join(opts, " ")

}

// CmdBuild run the build command.
func CmdBuild(c *cli.Context) error {

	opt := BuildOpt{
		Directories: Directories{
			Package:  c.GlobalString(PackageFlag),
			Homebrew: c.GlobalString(HomebrewFlag),
		},
		Version: c.Args().First(),
		GHROpt: ghrOpt{
			Token:   c.String("token"),
			Body:    c.String("body"),
			Process: c.Int("process"),
			Delete:  c.Bool("delete"),
			Draft:   c.Bool("draft"),
			Pre:     c.Bool("pre"),
		},
	}

	if err := cmdBuild(&opt); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil

}

func cmdBuild(opt *BuildOpt) (err error) {

	stdout := colorable.NewColorableStdout()
	stderr := colorable.NewColorableStderr()

	// Build and upload via make.
	fmt.Fprintln(stdout, chalk.Bold.TextStyle("Building binaries."))

	var cmd *exec.Cmd
	if opt.Version != "" {

		// If body is not given but CHANGELOG.md has a release note,
		// use it instead.
		if opt.GHROpt.Body == "" {
			var note string
			note, err = fgo.ReleaseNote("CHANGELOG.md", opt.Version)
			if err == nil {
				opt.GHROpt.Body = strings.Replace(note, `"`, `\"`, -1)
			}
		}

		ghrflags := fmt.Sprintf(`GHRFLAGS=%v`, opt.GHROpt.String())
		fmt.Println(ghrflags)
		cmd = exec.Command("make", "build", "release", fmt.Sprintf("VERSION=%s", opt.Version), ghrflags)
	} else {
		fmt.Println("Version is not given, set `snapshot`")
		cmd = exec.Command("make", "build")
	}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err = cmd.Run(); err != nil {
		return
	}

	return cmdUpdate(opt.Directories.Package, opt.Directories.Homebrew, opt.Version)

}
