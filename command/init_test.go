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
	"strings"
	"testing"
)

func TestCmdInit(t *testing.T) {

	// Move temporary directory.
	cd, temp, err := moveToTempDir(".", "test-cmd-init")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(temp)
	defer os.RemoveAll(temp)
	defer os.Chdir(cd)

	// Test w/ username.
	opt := InitOpt{
		Directories: Directories{
			Package:  "test-package",
			Homebrew: "test-homebrew",
		},
		UserName: "test-name",
	}
	if err = cmdInit(&opt); err != nil {
		t.Fatal(err.Error())
	}

	raw, err := ioutil.ReadFile("Makefile")
	if err != nil {
		t.Fatal(err.Error())
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
		t.Fatal(err.Error())
	}
	formula := string(raw)
	if !strings.Contains(formula, "https://github.com/test-name/") {
		t.Errorf("Formula template has wrong user name.\n%s\n", formula)
	}
	if !strings.Contains(formula, `desc ""`) {
		t.Error("Formula template has wrong description", formula)
	}
	os.Remove("Makefile")
	os.Remove("test-homebrew/fgo.rb.template")

	// Test w/ description.
	opt = InitOpt{
		Directories: Directories{
			Package:  "test-package",
			Homebrew: "test-homebrew",
		},
		UserName:    "test-name",
		Description: "sample description",
	}
	if err = cmdInit(&opt); err != nil {
		t.Fatal(err.Error())
	}

	raw, err = ioutil.ReadFile("test-homebrew/fgo.rb.template")
	if err != nil {
		t.Fatal(err.Error())
	}
	formula = string(raw)
	if !strings.Contains(formula, `desc "sample description"`) {
		t.Error("Formula template has wrong description", formula)
	}
	os.Remove("Makefile")
	os.Remove("test-homebrew/fgo.rb.template")

	// Test w/o username.
	// This test should be only run on local computers.
	if os.Getenv("LOCAL") == "true" {

		opt = InitOpt{
			Directories: Directories{
				Package:  "test-package",
				Homebrew: "test-homebrew",
			},
		}
		if err = cmdInit(&opt); err != nil {
			t.Fatal(err.Error())
		}
		raw, err = ioutil.ReadFile("Makefile")
		if err != nil {
			t.Fatal(err.Error())
		}
		makefile = string(raw)
		if !strings.Contains(makefile, "-u jkawamoto") {
			t.Errorf("Makefile has wrong user name.\n%s\n", makefile)
		}

	}

}

// TestPrepareDirectory tests prepareDirectory within the following three cases;
// 1) creating non existing directory, 2) using existing directory,
// 3) trying to make a directory of which name is colliding another file.
func TestPrepareDirectory(t *testing.T) {

	var err error

	// Move temporary directory.
	cd, temp, err := moveToTempDir("", "test-prepare-directory")
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer os.Chdir(cd)
	defer os.RemoveAll(temp)

	t.Log("Test with an existing directory.")
	target, err := ioutil.TempDir("", "fgo-test")
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer os.RemoveAll(target)

	if err = prepareDirectory(target); err != nil {
		t.Error(err.Error())
	}

	t.Log("Test with an existing file.")
	fp, err := ioutil.TempFile("", "fgo-test2")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fp.Close()

	if prepareDirectory(fp.Name()) == nil {
		t.Error("No error occurs when preparing directory with an existing file.")
	}

	if err = os.Remove(fp.Name()); err != nil {
		t.Error(err.Error())
		return
	}

	t.Log("Test without any collisions.")
	if err = prepareDirectory(fp.Name()); err != nil {
		t.Error(err.Error())
	}

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
