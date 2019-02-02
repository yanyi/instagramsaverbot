package scraper

import (
	"errors"
	"net/url"
	"strings"
)

var errNotInstagramLink = errors.New("Link given is not an Instagram link")

// getGraphURL appends "?__a=1" to back of Instagram URL after going through getRawURL().
func getGraphURL(inputURL string) (string, error) {
	instagramURL, err := urlIsInstagram(inputURL)
	if err != nil {
		return "", err
	}

	rawURL := getRawURL(instagramURL)
	var str strings.Builder
	str.WriteString(rawURL)
	str.WriteString("?__a=1")

	return str.String(), nil
}

// urlIsInstagram checks the input string given if it is an Instagram hostname.
// It gives a generic error message since it is expected to get just Instagram link.
// URLs like ":" will return the same generic error message.
func urlIsInstagram(inputURL string) (*url.URL, error) {
	u, err := url.Parse(inputURL)
	if err != nil {
		return nil, errNotInstagramLink
	}

	hostname := u.Hostname()
	if hostname != "instagram.com" && hostname != "www.instagram.com" {
		return nil, errNotInstagramLink
	}

	return u, nil
}

// getRawURL gets the raw URL without the '?'.
func getRawURL(instagramURL *url.URL) string {
	instagramURL.RawQuery = ""

	return instagramURL.String()
}
