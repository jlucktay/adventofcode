// Package fetchaocday will do a HTTP GET on the given date/year page of the Advent of Code challenge.
// The return value will be a Day struct.
package fetchaocday

import (
	"fmt"
	"log"
	"strconv"

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
			day.Description = s.Text()
		},
	)

	return *day
}
