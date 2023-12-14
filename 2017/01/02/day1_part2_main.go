// Package main for Advent of Code 2017, day 1, part 2
// http://adventofcode.com/2017/day/1
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
	fmt.Printf("Result for day 1, part 2: %d\n", decode(cleanInput))
}
