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
	// GitHub repository name.
	Repository string
}

// Generater is an interface provides Generate method.
type Generater interface {
	Generate() ([]byte, error)
}

// CmdInit parses options and run cmdInit.
func CmdInit(c *cli.Context) error {

	opt := InitOpt{
		Config: Config{
			Package:  c.GlobalString(PackageFlag),
			Homebrew: c.GlobalString(HomebrewFlag),
		},
		UserName:   c.Args().First(),
		Repository: c.Args().Get(1),
	}

	if err := cmdInit(&opt); err != nil {
		return cli.NewExitError(chalk.Red.Color(err.Error()), 1)
	}
	return nil

}

// cmdInit defines init command action.
func cmdInit(opt *InitOpt) (err error) {

	// Check user name.
	if opt.UserName == "" {
		fmt.Printf("Checking git configuration to get the user name: ")
		opt.UserName, err = gitconfig.GithubUser()
		if err != nil {
			return fmt.Errorf("Cannot find user name (%s)", err.Error())
		}
		fmt.Println(chalk.Yellow.Color(opt.UserName))
	}

	// Prepare directories.
	fmt.Printf("Preparing the directory to store a brew formula: ")
	if err = prepareDirectory(opt.Config.Homebrew); err != nil {
		return
	}
	fmt.Println("done")

	// Check Makefile doesn't exist and create it.
	fmt.Printf("Creating Makefile: ")
	err = createResource("Makefile", &Makefile{
		Dest:     opt.Config.Package,
		UserName: opt.UserName,
	})
	if err != nil {
		fmt.Printf(chalk.Yellow.Color("skipped (%s).\n"), err.Error())
	} else {
		fmt.Println("done")
	}

	// Check brew rb file doesn't exist and create it.
	fmt.Printf("Creating a template of homebrew formula: ")
	if opt.Repository == "" {
		opt.Repository, err = gitconfig.Repository()
	}
	if opt.Repository == "" {
		fmt.Printf(chalk.Red.Color("skipped (%s).\n"), err.Error())
		fmt.Println(chalk.Yellow.Color("You must re-run init command after setting a remote repository."))
	} else {
		err = createResource(filepath.Join(opt.Config.Homebrew, fmt.Sprintf("%s.rb.template", opt.Repository)), &FormulaTemplate{
			Package:  opt.Repository,
			UserName: opt.UserName,
		})
		if err != nil {
			fmt.Printf(chalk.Yellow.Color("skipped (%s).\n"), err.Error())
		} else {
			fmt.Println("done")
		}
	}

	fmt.Printf("Storing configurations to %s.\n", ConfigFile)
	return opt.Config.Save(ConfigFile)

}

// prepareDirectory creates a directory if necessary.
func prepareDirectory(path string) error {

	if info, exist := os.Stat(path); exist == nil && !info.IsDir() {
		return fmt.Errorf("cannot make directory %s", path)
	} else if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("cannot make directory %s (%s)", path, err.Error())
	}
	return nil

}

// createResource creates a resource from a given generator and stores it to
// a given path.
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
