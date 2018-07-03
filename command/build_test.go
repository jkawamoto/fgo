/*
 * build_test.go
 *
 * Copyright (c) 2016-2018 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package command

import (
	"regexp"
	"strings"
	"testing"
)

func TestGHROptString(t *testing.T) {

	var opt ghrOpt
	var flags string

	opt = ghrOpt{}
	flags = opt.String()
	if flags != "" {
		t.Error("Empty options don't return an empty string", flags)
	}

	opt = ghrOpt{
		Token: "test_token",
		Body:  "test_body",
	}
	flags = opt.String()
	if !strings.Contains(flags, "-t test_token") {
		t.Error("Generated flag string doesn't have a GitHub API token:", flags)
	}
	if !strings.Contains(flags, `-b "test_body"`) {
		t.Error("Generated flag string doesn't have a correct body:", flags)
	}
	if strings.Contains(flags, "-p") || strings.Contains(flags, "-delete") {
		t.Error("Generated flag string has a not given option:", flags)
	}

	opt = ghrOpt{
		Process: 10,
		Pre:     true,
	}
	flags = opt.String()
	if !regexp.MustCompile(`-p 10\s+`).MatchString(flags) {
		t.Error("Generated flag string doesn't specify the number of gorountines:", flags)
	}
	if !strings.Contains(flags, "-prerelease") {
		t.Error("Generated flags string doesn't have prerelase flag:", flags)
	}

}
