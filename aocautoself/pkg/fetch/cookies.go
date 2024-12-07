package fetch

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/browserutils/kooky"

	_ "github.com/browserutils/kooky/browser/firefox"
)

const (
	cookieDomain = "adventofcode.com"
	cookieName   = "session"
)

// FirefoxCookie reads the adventofcode.com session cookie from the default profile in the local install of Firefox.
func FirefoxCookie() (*http.Cookie, error) {
	session := ""

	for _, cookieStore := range kooky.FindAllCookieStores() {
		if !cookieStore.IsDefaultProfile() {
			continue
		}

		cookies, err := cookieStore.ReadCookies(kooky.Valid, kooky.DomainHasSuffix(cookieDomain), kooky.Name(cookieName))
		if err != nil {
			return nil, fmt.Errorf("reading cookies: %w", err)
		} else if len(cookies) != 1 {
			return nil, fmt.Errorf("wrong number of '%s' cookies from '%s' domain", cookieName, cookieDomain)
		}

		session = cookies[0].Value

		break
	}

	if session == "" {
		return nil, fmt.Errorf("could not find '%s' cookie from '%s' domain in default profile of Firefox cookie store",
			cookieName, cookieDomain)
	}

	if !regexp.MustCompile(`(?i)^[0-9a-f]{128}$`).MatchString(session) {
		return nil, fmt.Errorf("session cookie '%s' was not a 128 character hexadecimal", session)
	}

	cookie := &http.Cookie{
		Name:  cookieName,
		Value: session,
	}

	return cookie, nil
}
