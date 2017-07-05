//
// command/formula_test.go
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
		t.Error("Generated file has wrong class name.")
	}
	if !strings.Contains(res, "https://github.com/abcde/test") {
		t.Error("Generated file has wrong URL.")
	}
	if !strings.Contains(res, "bin.install \"test\"") {
		t.Error("Generated file has wrong install command.")
	}
	t.Log(res)

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
	}
	data, err = args.Generate(fp.Name())
	if err != nil {
		t.Error(err)
	}

	res := string(data)
	if !strings.Contains(res, "vtest.version/testname.mac.64") {
		t.Error("URL for 64bit mac is wrong:", res)
	}
	if !strings.Contains(res, "vtest.version/testname.mac.386") {
		t.Error("URL for 386 mac is wrong:", res)
	}
	if !strings.Contains(res, `sha256 "testhash.mac.64"`) {
		t.Error("Hash value for 64bit mac is wrong:", res)
	}
	if !strings.Contains(res, `sha256 "testhash.mac.386"`) {
		t.Error("Hash value for 386 mac is wrong:", res)
	}

}
