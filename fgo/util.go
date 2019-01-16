/*
 * util.go
 *
 * Copyright (c) 2016-2019 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package fgo

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/jkawamoto/fgo/fgo/assets"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func generateFromAsset(assetName string, param interface{}) (res []byte, err error) {

	fp, err := assets.Assets.Open(assetName)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(fp)
	if err != nil {
		return
	}
	return generate(data, param)

}

func generateFromFile(path string, param interface{}) (res []byte, err error) {

	fp, err := os.Open(path)
	if err != nil {
		return
	}
	defer fp.Close()

	data, err := ioutil.ReadAll(fp)
	if err != nil {
		return
	}

	return generate(data, param)

}

func generate(data []byte, param interface{}) (res []byte, err error) {

	funcMap := template.FuncMap{
		"Title":   strings.Title,
		"ToUpper": strings.ToUpper,
		"ToLower": strings.ToLower,
	}

	temp, err := template.New("").Funcs(funcMap).Parse(string(data))
	if err != nil {
		return
	}

	buf := bytes.Buffer{}
	if err = temp.Execute(&buf, &param); err != nil {
		return
	}

	res = buf.Bytes()
	return

}

// Sha256 computes sha256 hash of the given file.
func Sha256(path string) (res string, err error) {

	fp, err := os.Open(path)
	if err != nil {
		return
	}
	defer fp.Close()

	data, err := ioutil.ReadAll(fp)
	if err != nil {
		return
	}

	raw := sha256.Sum256(data)
	res = hex.EncodeToString(raw[:])
	return

}
