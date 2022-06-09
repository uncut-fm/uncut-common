package rss

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// parseDuration parses duration in format 00:00:00
func parseDuration(dur string) (int, error) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println("failed parsing duration "+dur+", err: ", err)
		}
	}()

	if dur == "" {
		return 0, nil
	}

	parts := strings.Split(dur, ":")

	var durationString string

	switch len(parts) {
	case 3:
		durationString = fmt.Sprintf("%sh%sm%ss", parts[0], parts[1], parts[2])
	case 2:
		durationString = fmt.Sprintf("%sm%ss", parts[0], parts[1])
	case 1:
		return strconv.Atoi(parts[0])
	}

	duration, err := time.ParseDuration(durationString)
	if err != nil {
		return 0, err
	}

	return int(duration.Seconds()), nil
}

func parseStringDescription(description string) string {
	if strings.Contains(description, "http") {
		description = wrapLinksWithATag(description)
	}
	description = strings.ReplaceAll(description, "\n", "<br/>")
	description = strings.ReplaceAll(description, "\r", "")

	return description
}

func wrapLinksWithATag(description string) string {
	m1 := regexp.MustCompile(`(http|https)\:\/\/(\S+)`)
	return m1.ReplaceAllString(description, `<a href="$1://$2" target="_blank" rel=“noreferrer">$1://$2</a>`)
}
