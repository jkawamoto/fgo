/*
 * update.go
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
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/jkawamoto/fgo/fgo"
	"github.com/mattn/go-colorable"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// CmdUpdate implements update command.
func CmdUpdate(c *cli.Context) error {
	stderr := colorable.NewColorableStderr()

	if c.NArg() != 1 {
		fmt.Fprintf(stderr, chalk.Red.Color("expected one argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	pkg := c.GlobalString(PackageFlag)
	brew := c.GlobalString(HomebrewFlag)

	if err := cmdUpdate(pkg, brew, c.Args().First()); err != nil {
		fmt.Fprintln(stderr, err)
		return cli.NewExitError("", 10)
	}
	return nil
}

// cmdUpdate retrieves archives of the specified version in the given directory
// pkg, and updates the brew formula in the given path brew.
// If version is empty, "snapshot" will be used instead.
func cmdUpdate(pkg, brew, version string) (err error) {

	stdout := colorable.NewColorableStdout()

	fmt.Fprintln(stdout, chalk.Bold.TextStyle("Updating brew formula."))
	if version == "" {
		version = SnapshotVersion
	} else {
		version = strings.TrimSuffix(version, "v")
	}

	param := fgo.Formula{
		Version: version,
	}

	var matches []string
	// Find archives for Mac.
	matches, err = filepath.Glob(filepath.Join(pkg, version, "*darwin*.zip"))
	if err != nil {
		return
	}
	for _, f := range matches {
		switch {
		case strings.Contains(f, "386"):
			param.Mac386.FileName = filepath.Base(f)
			param.Mac386.Hash, err = fgo.Sha256(f)
			if err != nil {
				return
			}

		case strings.Contains(f, "amd64"):
			param.Mac64.FileName = filepath.Base(f)
			param.Mac64.Hash, err = fgo.Sha256(f)
			if err != nil {
				return
			}
		}
	}

	// Find archives for Linux.
	matches, err = filepath.Glob(filepath.Join(pkg, version, "*linux*.tar.gz"))
	if err != nil {
		return
	}
	for _, f := range matches {
		switch {
		case strings.Contains(f, "386"):
			param.Linux386.FileName = filepath.Base(f)
			param.Linux386.Hash, err = fgo.Sha256(f)
			if err != nil {
				return
			}

		case strings.Contains(f, "amd64"):
			param.Linux64.FileName = filepath.Base(f)
			param.Linux64.Hash, err = fgo.Sha256(f)
			if err != nil {
				return
			}
		}
	}

	// Check binary files are found in local.
	if param.Mac386.FileName == "" || param.Mac64.FileName == "" {
		return fmt.Errorf(chalk.Red.Color("Binary files are not found. Run build command instead"))
	}

	matches, err = filepath.Glob(filepath.Join(brew, "*"+BrewFormulaSuffix))
	if err != nil {
		return
	}
	if len(matches) == 0 {
		return fmt.Errorf("no brew formula template exists")
	}

	data, err := param.Generate(matches[0])
	if err != nil {
		return
	}

	return ioutil.WriteFile(strings.TrimSuffix(matches[0], TemplateSuffix), data, 0644)

}
