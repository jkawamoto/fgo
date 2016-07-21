package command

import (
	"io/ioutil"

	"github.com/naoina/toml"
)

// ConfigFile defines configuration file name.
const ConfigFile = ".fgo"

// Config defines a configuration structure.
type Config struct {
	// Directory to store built packages.
	Package string `toml:"package"`
	// Directory to store homwbrew formula.
	Homebrew string `toml:"homebrew"`
}

// Load loads configurations from a given path.
func (c *Config) Load(path string) (err error) {

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	return toml.Unmarshal(buf, c)

}

// Save saves configurations to a given path.
func (c *Config) Save(path string) (err error) {

	data, err := toml.Marshal(*c)
	if err != nil {
		return
	}

	return ioutil.WriteFile(path, data, 0644)

}
