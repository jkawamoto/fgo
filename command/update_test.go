/*
 * update_test.go
 *
 * Copyright (c) 2016-2018 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package command

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jkawamoto/fgo/fgo"
)

const (
	TestPackageRoot = "../pkg"
)

func TestCmdUpdate(t *testing.T) {

	pkgs, err := filepath.Glob(filepath.Join(TestPackageRoot, "*"))
	if err != nil {
		t.Fatal("faild to find test packages:", err)
	}

	dir, err := ioutil.TempDir("", "fgo")
	if err != nil {
		t.Fatal("faild to create a temporary directory:", err)
	}
	defer os.RemoveAll(dir)

	temp := fgo.FormulaTemplate{
		Package:  "fgo",
		UserName: "test-user",
	}
	data, err := temp.Generate()
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(filepath.Join(dir, "fgo.rb.template"), data, 0644)
	if err != nil {
		t.Fatal(err)
	}

	for _, pkg := range pkgs {
		pkg = filepath.ToSlash(pkg)
		sp := strings.Split(pkg, "/")
		version := sp[len(sp)-1]

		err = cmdUpdate(TestPackageRoot, dir, version)
		if err != nil {
			t.Error(err)
		}

		data, err = ioutil.ReadFile(filepath.Join(dir, "fgo.rb"))
		if err != nil {
			t.Error(err)
		}
		res := string(data)

		for _, ext := range []string{"*.zip", "*.tar.gz"} {

			archives, err := filepath.Glob(filepath.Join(pkg, ext))
			if err != nil {
				t.Fatal(err)
			}
			for _, archive := range archives {
				if !strings.Contains(res, filepath.Base(archive)) && !strings.Contains(archive, "arm") {
					t.Error("Created formula doesn't contain file path to", archive)
				}
			}

		}

	}

}
