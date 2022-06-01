package model

import (
	"bytes"
	"html/template"
)

type EmailReceiver struct {
	Email string
	Name  string
}

// ParseTemplate parses an email template, and returns it as an HTML
func ParseTemplate(path string, data interface{}) (string, error) {
	t, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
