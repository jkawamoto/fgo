/*
 * dir.go
 *
 * Copyright (c) 2016-2018 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package command

const (
	// PackageFlag defines the flag name of package option.
	PackageFlag = "pkg"
	// HomebrewFlag defines the flag name of homebrew option.
	HomebrewFlag = "brew"
)

// Directories defines a configuration structure.
type Directories struct {
	// Directory to store built packages.
	Package string
	// Directory to store homebrew formula.
	Homebrew string
}
