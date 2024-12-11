package fetch

import (
	"context"
	"net/http"
)

// Input fetches the Advent of Code puzzle input for the given year and day.
func Input(ctx context.Context, session *http.Cookie, year, day int) (string, error) {
	details, err := getFromWeb(ctx, year, day, awrInput, session)
	if err != nil {
		return "", err
	}

	return string(details), nil
}
