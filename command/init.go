/*
 * init.go
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
	"os"
	"path/filepath"

	"github.com/deiwin/interact"
	"github.com/jkawamoto/fgo/fgo"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"

	"io"

	"github.com/tcnksm/go-gitconfig"
)

// InitOpt defines options for cmdInit.
type InitOpt struct {
	// Directory configurations.
	Directories
	// GitHub user name.
	UserName string
	// GitHub repository name.
	Repository string
	// Description of the target application.
	Description string
	// Command name.
	CmdName string
	// Writer to output messages.
	Stdout io.Writer
}

// Generator is an interface provides Generate method.
type Generator interface {
	Generate() ([]byte, error)
}

// CmdInit parses options and run cmdInit.
func CmdInit(c *cli.Context) error {

	opt := InitOpt{
		Directories: Directories{
			Package:  c.GlobalString(PackageFlag),
			Homebrew: c.GlobalString(HomebrewFlag),
		},
		UserName:    c.Args().First(),
		Repository:  c.Args().Get(1),
		Description: c.String("desc"),
		CmdName:     c.String("name"),
		Stdout:      c.App.Writer,
	}
	return cmdInit(&opt)

}

// cmdInit defines init command action.
func cmdInit(opt *InitOpt) (err error) {

	// Check if a user name is given.
	if opt.UserName == "" {
		//noinspection GoUnhandledErrorResult
		fmt.Fprintf(opt.Stdout, "Checking git configuration to get the user name: ")
		opt.UserName, err = gitconfig.GithubUser()
		if err != nil {
			return cli.NewExitError(fmt.Sprintf(chalk.Red.Color("Cannot find user name (%v)"), err), 1)
		}
		//noinspection GoUnhandledErrorResult
		fmt.Fprintln(opt.Stdout, chalk.Yellow.Color(opt.UserName))
	}

	// Create directories if not exist.

	//noinspection GoUnhandledErrorResult
	fmt.Fprintf(opt.Stdout, "Preparing the directory to store a brew formula: ")
	if err = prepareDirectory(opt.Directories.Homebrew); err != nil {
		return cli.NewExitError(fmt.Sprintf(chalk.Red.Color("failed (%v)"), err), 2)
	}
	//noinspection GoUnhandledErrorResult
	fmt.Fprintln(opt.Stdout, chalk.Green.Color("done"))

	// Check if Makefile exists and create it if necessary.
	actor := interact.NewActor(os.Stdin, opt.Stdout)

	createMakefile := true
	if _, exist := os.Stat("Makefile"); exist == nil {
		createMakefile, err = actor.Confirm("Makefile already exists. Would you like to overwrite it?", interact.ConfirmDefaultToNo)
		if err != nil {
			return cli.NewExitError(fmt.Sprintf(chalk.Red.Color("failed (%v)"), err), 3)
		}
	}
	if createMakefile {
		//noinspection GoUnhandledErrorResult
		fmt.Fprintf(opt.Stdout, "Creating Makefile: ")
		err = createResource("Makefile", &fgo.Makefile{
			Dest:     opt.Directories.Package,
			UserName: opt.UserName,
		})
		if err != nil {
			//noinspection GoUnhandledErrorResult
			fmt.Fprintf(opt.Stdout, chalk.Yellow.Color("skipped (%s)\n"), err)
		} else {
			//noinspection GoUnhandledErrorResult
			fmt.Fprintln(opt.Stdout, chalk.Green.Color("done"))
		}
	} else {
		//noinspection GoUnhandledErrorResult
		fmt.Fprintln(opt.Stdout, "Creating Makefile:", chalk.Yellow.Color("skipped"))
	}

	// Check if a template of Homebrew configuration file exists and create it if necessary.
	if opt.Repository == "" {
		//noinspection GoUnhandledErrorResult
		fmt.Fprintf(opt.Stdout, "Checking git configuration to get the repository name: ")
		opt.Repository, err = gitconfig.Repository()
		if err != nil {
			//noinspection GoUnhandledErrorResult
			fmt.Fprintf(opt.Stdout, chalk.Red.Color("skipped (%s).\n"), err)
			//noinspection GoUnhandledErrorResult
			fmt.Fprintln(opt.Stdout, chalk.Yellow.Color("You must re-run init command after setting a remote repository"))
		}
		//noinspection GoUnhandledErrorResult
		fmt.Fprintln(opt.Stdout, chalk.Yellow.Color(opt.Repository))
	}
	if opt.Repository != "" {

		if opt.CmdName == "" {
			opt.CmdName = opt.Repository
		}
		tmpFile := filepath.Join(opt.Directories.Homebrew, opt.CmdName+BrewFormulaSuffix)

		createTemplate := true
		if _, exist := os.Stat(tmpFile); exist == nil {
			createTemplate, err = actor.Confirm("brew formula template already exists. Would you like to overwrite it?", interact.ConfirmDefaultToNo)
			if err != nil {
				return cli.NewExitError(fmt.Sprintf(chalk.Red.Color("failed (%v)\n"), err), 3)
			}
		}
		if createTemplate {
			//noinspection GoUnhandledErrorResult
			fmt.Fprintf(opt.Stdout, "Creating brew formula template: ")
			err = createResource(tmpFile, &fgo.FormulaTemplate{
				Package:     opt.Repository,
				UserName:    opt.UserName,
				Description: opt.Description,
			})
			if err != nil {
				//noinspection GoUnhandledErrorResult
				fmt.Fprintf(opt.Stdout, chalk.Yellow.Color("skipped (%s)\n"), err)
			} else {
				//noinspection GoUnhandledErrorResult
				fmt.Fprintln(opt.Stdout, chalk.Green.Color("done"))
			}
		} else {
			//noinspection GoUnhandledErrorResult
			fmt.Fprintln(opt.Stdout, "Creating brew formula template:", chalk.Yellow.Color("skipped"))
		}
	}

	return

}

// prepareDirectory creates a directory if necessary.
func prepareDirectory(path string) error {

	if info, exist := os.Stat(path); exist == nil && !info.IsDir() {
		return fmt.Errorf("cannot make directory %s", path)
	} else if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("cannot make directory %s (%s)", path, err)
	}
	return nil

}

// createResource creates a resource from a given generator and stores it to
// a given path.
func createResource(path string, data Generator) (err error) {

	buf, err := data.Generate()
	if err != nil {
		return
	}
	return ioutil.WriteFile(path, buf, 0644)

}
