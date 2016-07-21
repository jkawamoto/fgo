//
// command/init_test.go
//
// Copyright (c) 2016 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

package command

import (
	"io/ioutil"
	"os"
	"testing"
)

// TestPrepareDirectory tests prepareDirectory within the following three cases;
// 1) creating non existing directory, 2) using existing directory,
// 3) trying to make a directory of which name is colliding another file.
func TestPrepareDirectory(t *testing.T) {

	var err error

	// Move temporary directory.
	cd, err := os.Getwd()
	if err != nil {
		t.Error(err.Error())
		return
	}
	if err = os.Chdir(os.TempDir()); err != nil {
		t.Error(err.Error())
		return
	}
	defer os.Chdir(cd)

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
