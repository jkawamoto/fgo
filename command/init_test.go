/*
 * init_test.go
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
	"strings"
	"testing"
)

func TestCmdInit(t *testing.T) {

	// Move temporary directory.
	cd, temp, err := moveToTempDir(".", "test-cmd-init")
	if err != nil {
		t.Fatal("failed to prepare a temporary directory:", err)
	}
	defer os.RemoveAll(temp)
	defer os.Chdir(cd)

	t.Run("with a user name", func(t *testing.T) {

		opt := InitOpt{
			Directories: Directories{
				Package:  "test-package",
				Homebrew: "test-homebrew",
			},
			UserName: "test-name",
			Stdout:   ioutil.Discard,
		}
		if err = cmdInit(&opt); err != nil {
			t.Fatal("cmdInit returned an error:", err)
		}
		defer os.Remove("Makefile")
		defer os.Remove("test-homebrew/fgo.rb.template")

		raw, err := ioutil.ReadFile("Makefile")
		if err != nil {
			t.Fatal("failed to read Makefile:", err)
		}
		makefile := string(raw)
		if !strings.Contains(makefile, "goxc -d=test-package") {
			t.Errorf("Makefile has wrong destination.\n%s\n", makefile)
		}
		if !strings.Contains(makefile, "-u test-name") {
			t.Errorf("Makefile has wrong user name.\n%s\n", makefile)
		}

		raw, err = ioutil.ReadFile("test-homebrew/fgo.rb.template")
		if err != nil {
			t.Fatal("failed to read the template", err)
		}
		formula := string(raw)
		if !strings.Contains(formula, "https://github.com/test-name/") {
			t.Errorf("Formula template has wrong user name.\n%s\n", formula)
		}
		if !strings.Contains(formula, `desc ""`) {
			t.Error("Formula template has wrong description", formula)
		}

	})

	t.Run("with description", func(t *testing.T) {

		opt := InitOpt{
			Directories: Directories{
				Package:  "test-package",
				Homebrew: "test-homebrew",
			},
			UserName:    "test-name",
			Description: "sample description",
			Stdout:      ioutil.Discard,
		}
		if err = cmdInit(&opt); err != nil {
			t.Fatal("cmdInit returned an error:", err)
		}
		defer os.Remove("Makefile")
		defer os.Remove("test-homebrew/fgo.rb.template")

		raw, err := ioutil.ReadFile("test-homebrew/fgo.rb.template")
		if err != nil {
			t.Fatal("failed to read the template file:", err)
		}
		formula := string(raw)
		if !strings.Contains(formula, `desc "sample description"`) {
			t.Error("Formula template has wrong description", formula)
		}

	})

	t.Run("with a command name", func(t *testing.T) {

		cmdName := "test-command"
		template := filepath.Join("test-homebrew", fmt.Sprintf("%v.rb.template", cmdName))
		opt := InitOpt{
			Directories: Directories{
				Package:  "test-package",
				Homebrew: "test-homebrew",
			},
			UserName: "test-name",
			CmdName:  cmdName,
			Stdout:   ioutil.Discard,
		}
		if err = cmdInit(&opt); err != nil {
			t.Fatal("cmdInit returned an error:", err)
		}
		defer os.Remove("Makefile")
		defer os.Remove(template)

		_, err := os.Stat(template)
		if err != nil {
			t.Error("the template file doesn't exists:", err)
		}

	})

	t.Run("without a user name", func(t *testing.T) {
		if os.Getenv("LOCAL") != "true" {
			t.Skip("This test should be only run on local computers")
		}

		opt := InitOpt{
			Directories: Directories{
				Package:  "test-package",
				Homebrew: "test-homebrew",
			},
			Stdout: ioutil.Discard,
		}
		if err = cmdInit(&opt); err != nil {
			t.Fatal("cmdInit returned an error:", err)
		}
		raw, err := ioutil.ReadFile("Makefile")
		if err != nil {
			t.Fatal("failed to read Makefile", err)
		}
		makefile := string(raw)
		if !strings.Contains(makefile, "-u jkawamoto") {
			t.Errorf("Makefile has wrong user name.\n%s\n", makefile)
		}

	})

}

// TestPrepareDirectory tests prepareDirectory within the following three cases;
// 1) creating non existing directory, 2) using existing directory,
// 3) trying to make a directory of which name is colliding another file.
func TestPrepareDirectory(t *testing.T) {

	// Move temporary directory.
	cd, temp, err := moveToTempDir("", "test-prepare-directory")
	if err != nil {
		t.Fatal("faild to prepare a temporary directory:", err)
	}
	defer os.Chdir(cd)
	defer os.RemoveAll(temp)

	t.Run("with an existing directory", func(t *testing.T) {

		target, err := ioutil.TempDir("", "fgo-test")
		if err != nil {
			t.Fatal("failed to prepare a temporary directory:", err)
		}
		defer os.RemoveAll(target)

		if err = prepareDirectory(target); err != nil {
			t.Error("failed to prepare the target directory:", err)
		}

	})

	t.Run("with an existing file", func(t *testing.T) {

		fp, err := ioutil.TempFile("", "fgo-test2")
		if err != nil {
			t.Fatal("failed to prepare a temporary directory:", err)
		}
		defer os.Remove(fp.Name())
		fp.Close()

		if prepareDirectory(fp.Name()) == nil {
			t.Error("No error occurs when preparing directory with an existing file.")
		}

	})

	t.Run("without any collisions", func(t *testing.T) {
		fp, err := ioutil.TempFile("", "fgo-test3")
		if err != nil {
			t.Fatal("failed to prepare a temporary directory:", err)
		}
		fp.Close()
		os.Remove(fp.Name())

		if err = prepareDirectory(fp.Name()); err != nil {
			t.Error(err.Error())
		}
	})

}

func moveToTempDir(dir, prefix string) (cd, temp string, err error) {

	// Prepare test directory.
	cd, err = os.Getwd()
	if err != nil {
		return
	}

	temp, err = ioutil.TempDir(dir, prefix)
	if err != nil {
		return
	}

	err = os.Chdir(temp)
	return

}
