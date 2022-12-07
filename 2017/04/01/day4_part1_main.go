// Package main for Advent of Code 2017, day 4, part 1
// http://adventofcode.com/2017/day/4
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var tally uint

	for _, p := range strings.Split(string(input[:]), "\n") {
		if validatePassphrase(p) {
			tally++
		}
	}

	fmt.Printf("Result for day 4, part 1: %d\n", tally)
}
