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

func CmdUpdate(c *cli.Context) error {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected one argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	pkg := c.GlobalString("dest")
	brew := c.GlobalString("brew")

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
			param.FileName386 = filepath.Base(f)
			param.Hash386, err = Sha256(f)
			if err != nil {
				return
			}

		case strings.Contains(f, "amd64"):
			param.FileName64 = filepath.Base(f)
			param.Hash64, err = Sha256(f)
			if err != nil {
				return
			}
		}
	}

	data, err := param.Generate(filepath.Join(brew, fmt.Sprintf("%s.rb.template", repo)))
	if err != nil {
		return
	}
	return ioutil.WriteFile(filepath.Join(brew, fmt.Sprintf("%s.rb", repo)), data, 0644)

}
