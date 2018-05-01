// Package fetchaocday will do a HTTP GET on the given date/year page of the Advent of Code challenge.
// The return value will be a Day struct.
package fetchaocday

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jlucktay/adventofcode/aocautoself/pkg"
	"github.com/jlucktay/adventofcode/aocautoself/pkg/aocday"
)

// Fetch will perform the necessary HTTP magic to get an Advent of Code day for the given year and date.
func Fetch(y, d uint) aocautoself.Day {
	day := aocday.NewDay(y, d)

	url := fmt.Sprintf("http://adventofcode.com/%d/day/%d", y, d)

	fmt.Println("[fetchaocday.Fetch] Fetching '" + url + "'...")

	doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("article.day-desc").Each(
		func(i int, s *goquery.Selection) {
			fmt.Println("[fetchaocday.Fetch] i: '" + strconv.Itoa(i) + "'")

			switch i {
			case 0:
				breakDescriptionDown(&day.Part1, s.Text(), i)
			case 1:
				breakDescriptionDown(&day.Part2, s.Text(), i)
			default:
				panic(fmt.Sprintf("don't know what to do with %d article.day-desc elements", i+1))
			}

			day.Description += s.Text()
		},
	)

	return *day
}

func breakDescriptionDown(dd *aocautoself.DayDesc, desc string, part int) {
	descSlice := strings.Split(desc, "\n")
	endsWithColon, endsWithQuestion := false, false

	for _, y := range descSlice {
		if len(y) > 0 {
			switch strings.TrimSpace(y)[len(y)-1:] {
			case ":":
				endsWithColon = true
			case "?":
				endsWithQuestion = true
			}
		}

		if !endsWithColon && !endsWithQuestion {
			dd.Fluff += y + "\n"
		} else if !endsWithQuestion {
			dd.Test += y + "\n"
		} else {
			dd.Stinger += y + "\n"
		}
	}

	dd.Fluff = strings.TrimSpace(dd.Fluff)
	dd.Test = strings.TrimSpace(dd.Test)
	dd.Stinger = strings.TrimSpace(dd.Stinger)
}
