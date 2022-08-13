package stringx

import (
	"bytes"
	"text/template"
)

// ParseTextTemplate parse data into the template
func ParseTextTemplate(tmplName, tmplConst string, data map[string]interface{}) (*bytes.Buffer, error) {
	tmp, err := template.New(tmplName).Parse(tmplConst)

	if err != nil {
		return nil, err
	}
	buff := new(bytes.Buffer)
	err = tmp.Execute(buff, data)
	if err != nil {
		return nil, err
	}

	return buff, nil
}
