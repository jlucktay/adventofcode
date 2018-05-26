// Package fetchaocday will do a HTTP GET on the given date/year page of the Advent of Code challenge.
// The return value will be a Day struct.
package fetchaocday

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jlucktay/adventofcode/aocautoself/pkg"
	"github.com/jlucktay/adventofcode/aocautoself/pkg/aocday"
)

// Fetch will perform the necessary HTTP magic to get an Advent of Code day for the given year and date.
func Fetch(cookie string, y, d uint) (output aocautoself.Day) {
	output = *aocday.NewDay(y, d)
	dayURL, _ := url.Parse(fmt.Sprintf("http://adventofcode.com/%d/day/%d", output.Year, output.Date))

	res, err := http.DefaultClient.Do(newRequest(*dayURL, cookie))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// fmt.Printf("[fetchaocday.Fetch] Fetching '%s'...\n", dayURL)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("article.day-desc").Each(
		func(i int, s *goquery.Selection) {
			// fmt.Println("[fetchaocday.Fetch] i: '" + strconv.Itoa(i) + "'")

			switch i {
			case 0:
				breakDescriptionDown(&output.Part1, s.Text())
			case 1:
				breakDescriptionDown(&output.Part2, s.Text())
			default:
				log.Fatalf("'%d' article.day-desc elements is too many", i+1)
			}

			output.Description += s.Text()
		},
	)

	return
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

func newRequest(u url.URL, sessionCookie string) (req *http.Request) {
	cookiePattern := `(?i)^[0-9a-f]{96,96}$`
	r, err := regexp.Compile(cookiePattern)
	if err != nil {
		log.Fatalf("regex '%s' couldn't compile: %s", cookiePattern, err)
	}

	if !r.MatchString(sessionCookie) {
		log.Fatalf("session cookie '%s' was not a 96 character hexadecimal", sessionCookie)
	}

	req, err = http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Cookie", fmt.Sprintf("session=%s", sessionCookie))

	return req
}
