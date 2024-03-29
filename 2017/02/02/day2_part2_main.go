// Package main for Advent of Code 2017, day 2, part 2
// http://adventofcode.com/2017/day/2
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	cleanInput := strings.TrimSpace(string(input))
	fmt.Printf("Result for day 2, part 2: %d\n", evenlyDivisibleChecksum(cleanInput))
}
