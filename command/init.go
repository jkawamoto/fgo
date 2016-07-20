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
	// Directory to store package files
	Dest string
	// Directory to store brew file
	Brew string
	// GitHub user name.
	UserName string
}

type Generater interface {
	Generate() ([]byte, error)
}

func CmdInit(c *cli.Context) error {

	opt := InitOpt{
		Dest:     c.String("dest"),
		Brew:     c.String("brew"),
		UserName: c.Args().First(),
	}

	// These codes are not nessesary but urfave/cli doesn't work.
	if opt.Dest == "" {
		opt.Dest = "pkg"
	}
	if opt.Brew == "" {
		opt.Brew = "brew"
	}

	if err := cmdInit(&opt); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil

}

func cmdInit(opt *InitOpt) (err error) {

	// Prepare directories.
	if err = prepareDirectory(opt.Dest); err != nil {
		return
	}
	if err = prepareDirectory(opt.Brew); err != nil {
		return
	}

	// Check Makefile doesn't exist and create it.
	err = createResource("Makefile", &Makefile{
		Dest:     opt.Dest,
		UserName: opt.UserName,
	})
	if err != nil {
		return
	}

	// Check brew rb file doesn't exist and create it.
	repo, err := gitconfig.Repository()
	if err != nil {
		return
	}
	username, err := gitconfig.Username()
	if err != nil {
		return
	}
	err = createResource(filepath.Join(opt.Brew, fmt.Sprintf("%s.rb", repo)), &Formula{
		Package:  repo,
		UserName: username,
	})
	if err != nil {
		return
	}

	return

}

// prepareDirectory creates a dicretory if necessary.
func prepareDirectory(path string) error {

	if info, exist := os.Stat(path); exist == nil && !info.IsDir() {
		return fmt.Errorf(chalk.Red.Color("Cannot make directory %s."), path)
	} else if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf(
			chalk.Red.Color("Cannot make directory %s. (%s)"),
			path, err.Error())
	}
	return nil

}

func createResource(path string, data Generater) (err error) {

	if _, exist := os.Stat(path); exist == nil {
		return fmt.Errorf(chalk.Red.Color("%s already exists."), path)
	}
	buf, err := data.Generate()
	if err != nil {
		return
	}
	return ioutil.WriteFile(path, buf, 0644)

}
