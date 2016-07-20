package command

import (
	"strings"
	"testing"
)

// TestMakefile tests generated Makefile contains a name given as a parameter.
func TestMakefile(t *testing.T) {

	param := Makefile{
		Dest: "test",
	}

	data, err := param.Generate()
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !strings.Contains(string(data), "-d=test") {
		t.Errorf("Generated Makefile was wrong.\n%s", string(data))
	}

}
