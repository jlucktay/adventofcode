package fetch

import (
	"fmt"
	"regexp"

	"github.com/zellyn/kooky"
	_ "github.com/zellyn/kooky/browser/firefox"
)

const (
	cookieDomain = "adventofcode.com"
	cookieName   = "session"
)

func SessionCookie() (string, error) {
	session := ""

	for _, cookieStore := range kooky.FindAllCookieStores() {
		if !cookieStore.IsDefaultProfile() {
			continue
		}

		cookies, err := cookieStore.ReadCookies(kooky.Valid, kooky.Domain(cookieDomain), kooky.Name(cookieName))
		if err != nil {
			return "", fmt.Errorf("reading cookies: %w", err)
		} else if len(cookies) != 1 {
			return "", fmt.Errorf("wrong number of '%s' cookies from '%s' domain", cookieName, cookieDomain)
		}

		session = cookies[0].Value
	}

	if session == "" {
		return "", fmt.Errorf("could not find '%s' cookie from '%s' domain in default profile of Firefox cookie store",
			cookieName, cookieDomain)
	}

	if !regexp.MustCompile(`(?i)^[0-9a-f]{128}$`).MatchString(session) {
		return "", fmt.Errorf("session cookie '%s' was not a 128 character hexadecimal", session)
	}

	return session, nil
}
