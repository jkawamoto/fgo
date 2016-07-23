package command

import (
	"fmt"
	"os"

	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

const (
	// DefaultPackageDir defines default package directory.
	DefaultPackageDir = "pkg"
	// DefaultHomebrewDir defines default homebrew formula directory.
	DefaultHomebrewDir = "homebrew"
	// PackageFlag defines the flag name of package option.
	PackageFlag = "pkg"
	// HomebrewFlag defines the flag name of homebrew option.
	HomebrewFlag = "brew"
)

// Prepare checkes configuration file and loads it if exists.
// Otherwise, set default values. In both cases, if optional flags are given,
// overwrite configurations by the given values.
func Prepare(c *cli.Context) error {

	config := Config{
		Package:  DefaultPackageDir,
		Homebrew: DefaultHomebrewDir,
	}

	// Check configuration file.
	if _, exist := os.Stat(ConfigFile); exist == nil {

		if err := config.Load(ConfigFile); err != nil {
			fmt.Printf(
				chalk.Red.Color("Cannot read configuration %s (%s)."),
				ConfigFile, err.Error())
		}

	}

	// If configurations are not given, set them.
	if !c.GlobalIsSet(PackageFlag) {
		c.GlobalSet(PackageFlag, config.Package)
	}
	if !c.GlobalIsSet(HomebrewFlag) {
		c.GlobalSet(HomebrewFlag, config.Homebrew)
	}
	return nil

}
