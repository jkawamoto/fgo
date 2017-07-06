//
// fgo/makefile.go
//
// Copyright (c) 2016-2017 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

package fgo

// MakefileAsset defines the asset name of Makefile.
const MakefileAsset = "assets/Makefile"

// Makefile defines variables to generate a Makefile.
type Makefile struct {
	// Directory to store package files
	Dest string
	// GitHub user name.
	UserName string
}

// Generate creates a Makefile by given variables.
func (m *Makefile) Generate() (res []byte, err error) {

	return generateFromAsset(MakefileAsset, m)

}
