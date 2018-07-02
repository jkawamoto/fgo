/*
 * release_test.go
 *
 * Copyright (c) 2016-2018 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package fgo

import (
	"strings"
	"testing"
)

func TestReleaseNote(t *testing.T) {

	var err error
	_, err = ReleaseNote("missing_file", "0.2.1")
	if err == nil {
		t.Error("Give a missing file but no error is returned.")
	}

	var note string
	note, err = ReleaseNote("../CHANGELOG.md", "0.1.0")
	if err != nil {
		t.Error(err.Error())
	} else if !strings.Contains(note, "Initial release") {
		t.Error("Returned release note is not correct:", note)
	}

	note, err = ReleaseNote("../CHANGELOG.md", "0.2.1")
	if err != nil {
		t.Error(err.Error())
	} else if !strings.HasPrefix(note, "<h3>Fixed</h3><ul><li>Problems of parsing global options.</li></ul>") {
		t.Error("Returned release note is not correct:", note)
	}

	note, err = ReleaseNote("../CHANGELOG.md", "1234.0.0")
	if err != nil {
		t.Error(err.Error())
	} else if note != "" {
		t.Error("Returned release note is not correct:", note)
	}

}
