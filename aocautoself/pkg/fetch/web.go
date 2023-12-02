package fetch

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func getFromWeb(year, date int, sessionCookie string) ([]byte, error) {
	dayURL, err := url.Parse(fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, date))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", dayURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
