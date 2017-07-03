//
// command/formula.go
//
// Copyright (c) 2016-2017 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

package command

// FormulaTemplateAsset defines the asset name of a formula template.
const FormulaTemplateAsset = "assets/formula.rb"

// FormulaTemplate defines variables to generate a template of
// a homebrew formula.
type FormulaTemplate struct {
	Package  string
	UserName string
}

// Formula defines variables to generate a homebrew formula.
type Formula struct {
	Version     string
	FileName64  string
	FileName386 string
	Hash64      string
	Hash386     string
}

// Generate creates a template of a homebrew formula by given variables.
func (f *FormulaTemplate) Generate() (res []byte, err error) {

	return generateFromAsset(FormulaTemplateAsset, f)

}

// Generate creates a homebrew formula by given variables.
func (f *Formula) Generate(path string) (ref []byte, err error) {

	return generateFromFile(path, f)

}
