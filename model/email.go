package model

import (
	"bytes"
	"html/template"
)

var (
	SpaceLinkPattern = "%s/user/my-shows/%s/space/%s" // {web-app}/my-shows/{show-slug}/space/{space_slug}
	ShowLinkPattern  = "%s/shows/%s"                  // {web-app}/shows/{show-slug}
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
