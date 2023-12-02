package fetch

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func getFromWeb(ctx context.Context, year, date int, session *http.Cookie) ([]byte, error) {
	dayURL, err := url.Parse(fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, date))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, dayURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if err := session.Valid(); err != nil {
		return nil, err
	}

	req.AddCookie(session)

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
