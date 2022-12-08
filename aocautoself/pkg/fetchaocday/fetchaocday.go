// Package fetchaocday will do a HTTP GET on the given date/year page of the Advent of Code challenge.
// The return value will be a Day struct.
package fetchaocday

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/sync/errgroup"

	aocautoself "go.jlucktay.dev/adventofcode/aocautoself/pkg"
	"go.jlucktay.dev/adventofcode/aocautoself/pkg/aocday"
)

// Fetch will perform the necessary HTTP magic to get an Advent of Code day for the given year and date.
func Fetch(cookie string, year, day uint) (aocautoself.Day, error) {
	output := *aocday.NewDay(year, day)

	dayURL, err := url.Parse(fmt.Sprintf("http://adventofcode.com/%d/day/%d", output.Year, output.Date))
	if err != nil {
		return aocautoself.Day{}, err
	}

	req, err := newRequest(*dayURL, cookie)
	if err != nil {
		return aocautoself.Day{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return aocautoself.Day{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return aocautoself.Day{}, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return aocautoself.Day{}, err
	}

	g := errgroup.Group{}
	g.SetLimit(1)

	doc.Find("article.day-desc").Each(
		func(i int, s *goquery.Selection) {
			g.Go(func() error {
				switch i {
				case 0:
					breakDescriptionDown(&output.Part1, s.Text())
				case 1:
					breakDescriptionDown(&output.Part2, s.Text())
				default:
					return fmt.Errorf("'%d' article.day-desc elements is too many", i+1)
				}

				output.Description += s.Text()

				return nil
			})
		},
	)

	if err := g.Wait(); err != nil {
		return aocautoself.Day{}, fmt.Errorf("parsing doc: %w", err)
	}

	return output, nil
}

func breakDescriptionDown(dd *aocautoself.DayDesc, desc string) {
	descSlice := strings.Split(desc, "\n")
	endsWithColon, endsWithQuestion := false, false

	for _, line := range descSlice {
		if len(line) > 0 {
			switch strings.TrimSpace(line)[len(line)-1:] {
			case ":":
				endsWithColon = true
			case "?":
				endsWithQuestion = true
			}
		}

		if !endsWithColon && !endsWithQuestion {
			dd.Fluff += line + "\n"
		} else if !endsWithQuestion {
			dd.Test += line + "\n"
		} else {
			dd.Stinger += line + "\n"
		}
	}

	dd.Fluff = strings.TrimSpace(dd.Fluff)
	dd.Test = strings.TrimSpace(dd.Test)
	dd.Stinger = strings.TrimSpace(dd.Stinger)
}

func newRequest(u url.URL, sessionCookie string) (*http.Request, error) {
	cookiePattern := `(?i)^[0-9a-f]{128}$`
	r, err := regexp.Compile(cookiePattern)
	if err != nil {
		return nil, fmt.Errorf("regex '%s' couldn't compile: %w", cookiePattern, err)
	}

	if !r.MatchString(sessionCookie) {
		return nil, fmt.Errorf("session cookie '%s' was not a 128 character hexadecimal", sessionCookie)
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Cookie", fmt.Sprintf("session=%s", sessionCookie))

	return req, nil
}
