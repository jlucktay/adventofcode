// This app will create some Advent of Code templates for you.
package main

import (
	"fmt"
	"log"

	"github.com/jlucktay/adventofcode/aocautoself/pkg"
	"github.com/jlucktay/adventofcode/aocautoself/pkg/cookiemonster"
	"github.com/jlucktay/adventofcode/aocautoself/pkg/fetchaocday"
)

func main() {
	cookie := cookiemonster.GetCookieWithKey("adventofcode.com", "session")
	if cookie == "" {
		log.Fatal("'cookie' was empty")
	}

	days := make([]aocautoself.Day, 0)
	days = append(days, fetchaocday.Fetch(cookie, 2017, 1))
	fmt.Println(days[0].String())
}
