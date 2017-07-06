//
// command/config.go
//
// Copyright (c) 2016-2017 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

package command

const (
	// PackageFlag defines the flag name of package option.
	PackageFlag = "pkg"
	// HomebrewFlag defines the flag name of homebrew option.
	HomebrewFlag = "brew"
)

// Config defines a configuration structure.
type Config struct {
	// Directory to store built packages.
	Package string
	// Directory to store homwbrew formula.
	Homebrew string
}
