package model

import (
	"bytes"
	"fmt"
	"html/template"
)

var (
	SpaceLinkPattern = "%s/user/my-shows/%s/space/%s" // {web-app}/my-shows/{show-slug}/space/{space_slug}
	ShowLinkPattern  = "%s/show/%s"                   // {web-app}/shows/{show-slug}
	NftLinkPattern   = "%s/show/%s/nft/%d"            // {web-app}/shows/{show-slug}/nft/{nft_id}
)

func GetShowLink(webAppURL, showSlug string) string {
	return fmt.Sprintf(ShowLinkPattern, webAppURL, showSlug)
}

func GetSpaceLink(webAppURL, showSlug, spaceSlug string) string {
	return fmt.Sprintf(SpaceLinkPattern, webAppURL, showSlug, spaceSlug)
}

func GetNftLink(webAppURL, showSlug string, nftID int) string {
	return fmt.Sprintf(NftLinkPattern, webAppURL, showSlug, nftID)
}

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
