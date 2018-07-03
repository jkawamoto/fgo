/*
 * formula_test.go
 *
 * Copyright (c) 2016-2018 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

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
		Package:     "test",
		UserName:    "abcde",
		Description: "sample text",
	}

	data, err := param.Generate()
	if err != nil {
		t.Fatal("Generate returned an error:", err)
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
	if !strings.Contains(res, `desc "sample text"`) {
		t.Error("Generated file doesn't have correct description", res)
	}

}

func TestFormula(t *testing.T) {

	packageInfo := FormulaTemplate{
		Package:  "test",
		UserName: "test-user",
	}
	data, err := packageInfo.Generate()
	if err != nil {
		t.Fatal("Generate returned an error:", err)
	}

	fp, err := ioutil.TempFile("", "test-formula")
	if err != nil {
		t.Fatal("failed to create a temporary file:", err)
	}
	defer os.Remove(fp.Name())

	_, err = fp.Write(data)
	if err != nil {
		t.Fatal("failed to write a brew formula:", err)
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
		t.Fatal("Generate returned an error:", err)
	}

	res := string(data)
	for _, osType := range []string{"mac", "linux"} {
		for _, arch := range []string{"64", "386"} {
			if !strings.Contains(res, fmt.Sprintf("vtest.version/testname.%v.%v", osType, arch)) {
				t.Errorf("URL for %v-%v is wrong: %v", osType, arch, res)
			}
			if !strings.Contains(res, fmt.Sprintf(`sha256 "testhash.%v.%v"`, osType, arch)) {
				t.Errorf("Hash value for %v-%v is wrong: %v", osType, arch, res)
			}
		}
	}

}
