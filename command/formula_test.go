//
// command/formula_test.go
//
// Copyright (c) 2016 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

package command

import (
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
