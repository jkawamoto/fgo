/*
 * release.go
 *
 * Copyright (c) 2016-2018 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package fgo

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/russross/blackfriday"
)

// ReleaseNote reads the given named file and returns a release note for the
// given version. The given named file must be written in Markdown and have
// headers containing a version number.
func ReleaseNote(filename, version string) (note string, err error) {

	fp, err := os.Open(filename)
	if err != nil {
		return
	}
	defer fp.Close()

	exp := regexp.MustCompile(fmt.Sprintf(`^(#+)\s*%v.*$`, version))
	s := bufio.NewScanner(fp)
	var end *regexp.Regexp
	var body []string
	for s.Scan() {

		if end != nil {
			if end.MatchString(s.Text()) {
				break
			}
			body = append(body, s.Text())

		} else if m := exp.FindStringSubmatch(s.Text()); m != nil {
			end = regexp.MustCompile(fmt.Sprintf(`^%v[^#]+.*$`, m[1]))
		}

	}

	note = strings.Replace(
		string(blackfriday.Run([]byte(strings.Join(body, "\n")))),
		"\n", "", -1)
	return

}
