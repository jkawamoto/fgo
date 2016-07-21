//
// command/init.go
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
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ttacon/chalk"
	"github.com/urfave/cli"

	"github.com/tcnksm/go-gitconfig"
)

// InitOpt defines options for cmdInit.
type InitOpt struct {
	// Configuration
	Config Config
	// GitHub user name.
	UserName string
}

type Generater interface {
	Generate() ([]byte, error)
}

// CmdInit parses options and run cmdInit.
func CmdInit(c *cli.Context) error {

	opt := InitOpt{
		Config: Config{
			Package:  c.String("dest"),
			Homebrew: c.String("brew"),
		},
		UserName: c.Args().First(),
	}

	if err := cmdInit(&opt); err != nil {
		return cli.NewExitError(chalk.Red.Color(err.Error()), 1)
	}
	return nil

}

// cmdInit defines init command action.
func cmdInit(opt *InitOpt) (err error) {

	if opt.UserName == "" {
		opt.UserName, err = gitconfig.GithubUser()
		if err != nil {
			return fmt.Errorf("Cannot find user name (%s)", err.Error())
		}
	}

	// Prepare directories.
	if err = prepareDirectory(opt.Config.Homebrew); err != nil {
		return
	}

	// Check Makefile doesn't exist and create it.
	fmt.Println(chalk.Bold.TextStyle("Creating Makefile."))
	err = createResource("Makefile", &Makefile{
		Dest:     opt.Config.Package,
		UserName: opt.UserName,
	})
	if err != nil {
		fmt.Printf(chalk.Yellow.Color("Cannot create Makefile (%s).\n"), err.Error())
	}

	// Check brew rb file doesn't exist and create it.
	repo, err := gitconfig.Repository()
	if err != nil {
		return
	}
	fmt.Println(chalk.Bold.TextStyle("Creating a template of homebrew formula."))
	err = createResource(filepath.Join(opt.Config.Homebrew, fmt.Sprintf("%s.rb.template", repo)), &FormulaTemplate{
		Package:  repo,
		UserName: opt.UserName,
	})
	if err != nil {
		fmt.Printf(chalk.Yellow.Color("Cannot create a formula template (%s).\n"), err.Error())
	}

	fmt.Printf(chalk.Bold.TextStyle("Storing configurations to %s.\n"), ConfigFile)
	return opt.Config.Save(ConfigFile)

}

// prepareDirectory creates a directory if necessary.
func prepareDirectory(path string) error {

	if info, exist := os.Stat(path); exist == nil && !info.IsDir() {
		return fmt.Errorf("Cannot make directory %s", path)
	} else if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("Cannot make directory %s (%s)", path, err.Error())
	}
	return nil

}

func createResource(path string, data Generater) (err error) {

	if _, exist := os.Stat(path); exist == nil {
		return fmt.Errorf("%s already exists", path)
	}
	buf, err := data.Generate()
	if err != nil {
		return
	}
	return ioutil.WriteFile(path, buf, 0644)

}
