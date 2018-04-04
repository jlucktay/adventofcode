// Package autosquared is code that writes boilerplate code.
// It is a self-perpetuating package that gathers stories from the Advent of Code website and scaffolds out code
// skeletons for the given year and day.
package autosquared

import (
	"fmt"
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func do() {
	descriptions := make(map[int]map[string]string)

	getAdventDayDescriptions(1, 25, descriptions)

	for a, b := range descriptions {
		// fmt.Printf("a: '%v'\n", a)

		for c, d := range b {
			// fmt.Printf("c: '%v'\n", c)
			// fmt.Printf("d: '%v'\n", d)

			result, _ := regexp.MatchString("For example", d)

			if result {
				fmt.Printf("Day %02d (url: '%40s') has a \"For example\"!\n", a, c)
			} else {
				fmt.Printf("Day %02d (url: '%40s') did not have a \"For example\". :(\n", a, c)
			}
		}
	}
}

func getAdventDayDescriptions(firstDay, lastDay int, m map[int]map[string]string) {
	for index := firstDay; index <= lastDay; index++ {
		url := fmt.Sprintf("http://adventofcode.com/2017/day/%d", index)

		// fmt.Println("Fetching '" + url + "'...")

		doc, err := goquery.NewDocument(url)

		if err != nil {
			log.Fatal(err)
		}

		doc.Find("article.day-desc").Each(
			func(i int, s *goquery.Selection) {
				m[index] = make(map[string]string)
				m[index][url] = s.Text()
			},
		)
	}
}
