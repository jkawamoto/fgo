/*
 * formula.go
 *
 * Copyright (c) 2016-2018 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package fgo

// FormulaTemplateAsset defines the asset name of a formula template.
const FormulaTemplateAsset = "/formula.rb"

// FormulaTemplate defines variables to generate a template of
// a homebrew formula.
type FormulaTemplate struct {
	// Package name
	Package string
	// User name
	UserName string
	// Description of the package
	Description string
}

// ArchiveInfo defines information of an archive file.
type ArchiveInfo struct {
	// File name of the archive.
	FileName string
	// Hash value of the archive file.
	Hash string
}

// Formula defines variables to generate a homebrew formula.
type Formula struct {
	// Version.
	Version string
	// Archive information for 64bit mac
	Mac64 ArchiveInfo
	// Archive information for 386 mac
	Mac386 ArchiveInfo
	// Archive information for 64bit Linux
	Linux64 ArchiveInfo
	// Archive information for 386 Linux
	Linux386 ArchiveInfo
}

// Generate creates a template of a homebrew formula by given variables.
func (f *FormulaTemplate) Generate() (res []byte, err error) {

	return generateFromAsset(FormulaTemplateAsset, f)

}

// Generate creates a homebrew formula by given variables.
func (f *Formula) Generate(path string) (ref []byte, err error) {

	return generateFromFile(path, f)

}
