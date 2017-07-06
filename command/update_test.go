//
// command/update_test.go
//
// Copyright (c) 2016-2017 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

package command

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	TestPackageRoot = "../pkg"
)

func TestCmdUpdate(t *testing.T) {

	pkgs, err := filepath.Glob(filepath.Join(TestPackageRoot, "*"))
	if err != nil {
		t.Fatal(err)
	}

	dir, err := ioutil.TempDir("", "fgo")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	temp := FormulaTemplate{
		Package:  "fgo",
		UserName: "testuser",
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
