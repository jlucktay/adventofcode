package fetch

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func getFromWeb(ctx context.Context, year, date int, awr aocWebResource, session *http.Cookie) ([]byte, error) {
	rawURL := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, date)

	switch awr {
	case awrDay:
		// no-op; the base rawURL is already sufficient

	case awrInput:
		rawURL += "/input"

	default:
		if aocWebResources.Contains(awr) {
			return nil, fmt.Errorf("enum '%v' should be handled by fetch.getFromWeb", awr)
		}

		return nil, fmt.Errorf("unknown AoC web resource '%v'", awr)
	}

	dayURL, err := url.Parse(rawURL)
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

	hc := http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK status error GETing '%s': %s", dayURL, resp.Status)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
