//
// fgo/formula_test.go
//
// Copyright (c) 2016-2017 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

package fgo

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestFormulaTemplate(t *testing.T) {

	param := FormulaTemplate{
		Package:  "test",
		UserName: "abcde",
	}

	data, err := param.Generate()
	if err != nil {
		t.Error(err.Error())
		return
	}

	res := string(data)
	if !strings.Contains(res, "class Test") {
		t.Error("Generated file has wrong class name.", res)
	}
	if !strings.Contains(res, "https://github.com/abcde/test") {
		t.Error("Generated file has wrong URL.", res)
	}
	if !strings.Contains(res, "bin.install \"test\"") {
		t.Error("Generated file has wrong install command.", res)
	}

}

func TestFormula(t *testing.T) {

	packageInfo := FormulaTemplate{
		Package:  "test",
		UserName: "test-user",
	}
	data, err := packageInfo.Generate()
	if err != nil {
		t.Fatal(err)
	}

	fp, err := ioutil.TempFile("", "test-formula")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(fp.Name())

	_, err = fp.Write(data)
	if err != nil {
		t.Fatal(err)
	}
	fp.Close()

	args := Formula{
		Version: "test.version",
		Mac64: ArchiveInfo{
			FileName: "testname.mac.64",
			Hash:     "testhash.mac.64",
		},
		Mac386: ArchiveInfo{
			FileName: "testname.mac.386",
			Hash:     "testhash.mac.386",
		},
		Linux64: ArchiveInfo{
			FileName: "testname.linux.64",
			Hash:     "testhash.linux.64",
		},
		Linux386: ArchiveInfo{
			FileName: "testname.linux.386",
			Hash:     "testhash.linux.386",
		},
	}
	data, err = args.Generate(fp.Name())
	if err != nil {
		t.Error(err)
	}

	res := string(data)
	for _, os := range []string{"mac", "linux"} {
		for _, arch := range []string{"64", "386"} {
			if !strings.Contains(res, fmt.Sprintf("vtest.version/testname.%v.%v", os, arch)) {
				t.Errorf("URL for %v-%v is wrong: %v", os, arch, res)
			}
			if !strings.Contains(res, fmt.Sprintf(`sha256 "testhash.%v.%v"`, os, arch)) {
				t.Errorf("Hash value for %v-%v is wrong: %v", os, arch, res)
			}
		}
	}

}
