// Package main for Advent of Code 2017, day 5, part 1
// http://adventofcode.com/2017/day/5
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	s := strings.Split(strings.TrimSpace(string(input[:])), "\n")
	jumps := make([]int, len(s))

	for i := range s {
		jumps[i], _ = strconv.Atoi(s[i])
	}

	result := followJumpOffsets(jumps)

	fmt.Printf("Result for day 5, part 1: %d\n", result)
}
