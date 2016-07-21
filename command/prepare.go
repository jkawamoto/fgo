package command

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/naoina/toml"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

const (
	// ConfigFile defines configuration file name.
	ConfigFile = ".fgo"
	// DefaultPackageDir defines default package directory.
	DefaultPackageDir = "pkg"
	// DefaultHomebrewDir defines default homebrew formula directory.
	DefaultHomebrewDir = "homebrew"
)

// Config defines a configuration structure.
type Config struct {
	// Directory to store built packages.
	Package string
	// Directory to store homwbrew formula.
	Homebrew string
}

// Load loads configurations from a given path.
func (c *Config) Load(path string) (err error) {

	fp, err := os.Open(path)
	if err != nil {
		return
	}
	defer fp.Close()

	buf, err := ioutil.ReadAll(fp)
	if err != nil {
		return
	}

	return toml.Unmarshal(buf, &c)

}

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
	if !c.IsSet("pkg") {
		c.Set("pkg", config.Package)
	}
	if !c.IsSet("brew") {
		c.Set("brew", config.Homebrew)
	}
	return nil

}
