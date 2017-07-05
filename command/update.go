//
// command/update.go
//
// Copyright (c) 2016-2017 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

package command

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/tcnksm/go-gitconfig"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// CmdUpdate implements update command.
func CmdUpdate(c *cli.Context) error {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected one argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	pkg := c.GlobalString(PackageFlag)
	brew := c.GlobalString(HomebrewFlag)

	if err := cmdUpdate(pkg, brew, c.Args().First()); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}

func cmdUpdate(pkg, brew, version string) (err error) {

	fmt.Println(chalk.Bold.TextStyle("Updating brew formula."))
	repo, err := gitconfig.Repository()
	if err != nil {
		return
	}

	if version == "" {
		version = "snapshot"
	}

	param := Formula{
		Version: version,
	}

	glob := filepath.Join(pkg, version, "*darwin*.zip")
	matches, err := filepath.Glob(glob)
	if err != nil {
		return
	}
	for _, f := range matches {
		switch {
		case strings.Contains(f, "386"):
			param.Mac386.FileName = filepath.Base(f)
			param.Mac386.Hash, err = Sha256(f)
			if err != nil {
				return
			}

		case strings.Contains(f, "amd64"):
			param.Mac64.FileName = filepath.Base(f)
			param.Mac64.Hash, err = Sha256(f)
			if err != nil {
				return
			}
		}
	}

	// Check binary files are found in local.
	if param.Mac386.FileName == "" || param.Mac64.FileName == "" {
		return fmt.Errorf(chalk.Red.Color("Binary files are not found. Run build command instead.\n"))
	}

	data, err := param.Generate(filepath.Join(brew, fmt.Sprintf("%s.rb.template", repo)))
	if err != nil {
		return
	}
	return ioutil.WriteFile(filepath.Join(brew, fmt.Sprintf("%s.rb", repo)), data, 0644)

}
