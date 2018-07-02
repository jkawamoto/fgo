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

	"io"

	"github.com/jkawamoto/fgo/fgo"
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
	// Writer to output messages
	Stdout io.Writer
	// Writer to output error messages
	Stderr io.Writer
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
		Stdout: c.App.Writer,
		Stderr: c.App.ErrWriter,
	}

	if err := cmdBuild(&opt); err != nil {
		return cli.NewExitError(err, 1)
	}
	return nil

}

func cmdBuild(opt *BuildOpt) (err error) {

	// Build and upload via make.
	fmt.Fprintln(opt.Stdout, chalk.Bold.TextStyle("Building binaries."))

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
	cmd.Stdout = opt.Stdout
	cmd.Stderr = opt.Stderr

	if err = cmd.Run(); err != nil {
		return
	}

	return cmdUpdate(opt.Directories.Package, opt.Directories.Homebrew, opt.Version, opt.Stdout)

}
