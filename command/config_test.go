package command

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// TestConfig tests save and load functions of Config.
func TestConfig(t *testing.T) {

	// Prepare test directory.
	cd, err := os.Getwd()
	if err != nil {
		t.Error(err.Error())
		return
	}

	temp, err := ioutil.TempDir("", "config-test")
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer os.RemoveAll(temp)

	if err = os.Chdir(temp); err != nil {
		t.Error(err.Error())
		return
	}
	defer os.Chdir(cd)

	// Test saving.
	config := Config{
		Package:  "test-package",
		Homebrew: "test-homebrew",
	}
	if err = config.Save(ConfigFile); err != nil {
		t.Error(err.Error())
		return
	}

	raw, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := string(raw)
	if !strings.Contains(data, "package=\"test-package\"") {
		t.Errorf("Package information isn't saved.\n%s", data)
	}
	if !strings.Contains(data, "homebrew=\"test-homebrew\"") {
		t.Errorf("Package information isn't saved.\n%s", data)
	}

	// Test loading.
	var another Config
	if err = another.Load(ConfigFile); err != nil {
		t.Error(err.Error())
		return
	}
	if another.Package != "test-package" || another.Homebrew != "test-homebrew" {
		t.Errorf("Package information isn't loaded.\n%s", another)
	}

}
