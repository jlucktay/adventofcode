// This app will create some Advent of Code templates for you.
package main

import (
	"fmt"
	"log"

	"go.jlucktay.dev/adventofcode/aocautoself/pkg/fetch"
)

func main() {
	session, err := fetch.SessionCookie()
	if err != nil {
		log.Fatal(err)
	}

	day, err := fetch.Day(session, 2017, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(day.String())
}
