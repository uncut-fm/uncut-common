package model

import (
	"bytes"
	"fmt"
	"github.com/uncut-fm/uncut-common/pkg/config"
	"html/template"
	"regexp"
	"strconv"
	"strings"
)

var (
	baseURL                 = "https://uncut.network"
	showLinkPattern         = "https://%s.uncut.network" // https://{show_slug}.uncut.network
	feedLinkPattern         = "%s/feed"                  // {show_link}/feed
	collectiveLinkPattern   = "%s/collective"            // {show_link}/collective
	conversationLinkPattern = "%s/%d"                    // {feed_link}/feed/{conversation_id}
	nftLinkPattern          = "%s/nft/%d"                // {show_link}/nft/{nft_id}
	personalNftLinkPattern  = "%s/unft/%d"               // {show_link}/unft/{nft_id}

	userLinkPattern = "%s/user/%d" // {base_url}/user/{user_id}
)

func GetShowLink(environment, showSlug string) string {
	switch environment {
	case config.DevEnvironment, config.StageEnvironment:
		if len(showSlug) == 0 {
			return fmt.Sprintf(showLinkPattern, environment)
		}
		return fmt.Sprintf(showLinkPattern, fmt.Sprintf("%s.%s", showSlug, environment))
	case config.TestEnvironment:
		if len(showSlug) == 0 {
			return fmt.Sprintf(showLinkPattern, "qa")
		}
		return fmt.Sprintf(showLinkPattern, fmt.Sprintf("%s.%s", showSlug, "qa"))
	default:
		if len(showSlug) == 0 {
			return baseURL
		}
		return fmt.Sprintf(showLinkPattern, showSlug)
	}
}

func GetFeedLink(environment, showSlug string) string {
	showLink := GetShowLink(environment, showSlug)
	return fmt.Sprintf(feedLinkPattern, showLink)
}

func GetConversationLink(environment, showSlug string, conversationID int) string {
	feedLink := GetFeedLink(environment, showSlug)
	return fmt.Sprintf(conversationLinkPattern, feedLink, conversationID)
}

func GetCollectiveLink(environment, showSlug string) string {
	showLink := GetShowLink(environment, showSlug)
	return fmt.Sprintf(collectiveLinkPattern, showLink)
}

func GetNftLink(environment, showSlug string, nftID int) string {
	showLink := GetShowLink(environment, showSlug)
	if len(showSlug) == 0 {
		return fmt.Sprintf(personalNftLinkPattern, showLink, nftID)
	}

	return fmt.Sprintf(nftLinkPattern, showLink, nftID)
}

func GetIDFromNftLink(nftLink string) (int, error) {
	var pattern *regexp.Regexp
	if strings.Contains(nftLink, "/nft/") {
		pattern = regexp.MustCompile(`/nft/(\d+)$`)
	} else if strings.Contains(nftLink, "/unft/") {
		pattern = regexp.MustCompile(`/unft/(\d+)$`)
	} else {
		return 0, fmt.Errorf("unsupported URL pattern: %s", nftLink)
	}

	match := pattern.FindStringSubmatch(nftLink)
	if len(match) != 2 {
		return 0, fmt.Errorf("invalid URL format: %s", nftLink)
	}

	idString := match[1]

	return strconv.Atoi(idString)
}

func GetUserNetworkLink(environment string, userID int) string {
	return fmt.Sprintf("%s/network", GetUserLink(environment, userID))
}

func GetUserLink(environment string, userID int) string {
	baseLink := GetShowLink(environment, "")

	return fmt.Sprintf(userLinkPattern, baseLink, userID)
}

type EmailReceiver struct {
	Email string
	Name  string
	ID    int
}

// ParseTemplate parses an email template, and returns it as an HTML
func ParseTemplate(data interface{}, paths ...string) (string, error) {
	t, err := template.ParseFiles(paths...)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if len(paths) > 1 {
		if err = t.ExecuteTemplate(buf, "layout", data); err != nil {
			return "", err
		}
	} else {
		if err = t.Execute(buf, data); err != nil {
			return "", err
		}
	}

	return buf.String(), nil
}
