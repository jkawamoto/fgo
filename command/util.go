package command

import (
	"bytes"
	"strings"
	"text/template"
)

func generate(assetName string, param interface{}) (res []byte, err error) {

	funcMap := template.FuncMap{
		"Title":   strings.Title,
		"ToUpper": strings.ToUpper,
		"ToLower": strings.ToLower,
	}

	data, err := Asset(assetName)
	if err != nil {
		return
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
