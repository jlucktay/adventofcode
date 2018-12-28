package aoc201801

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// GetInput returns the content of 'input.txt' as a string
func GetInput() string {
	rawInput, err := ioutil.ReadFile("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	return string(rawInput)
}

// ProcessInput converts an input file string from newline-delimited integers to
// a slice of positive and negative ints
func ProcessInput(input string) []int {
	result := make([]int, 0)

	for _, i := range strings.Split(string(input), "\n") {
		convInt, convErr := strconv.Atoi(strings.TrimSpace(string(i)))

		if convErr != nil {
			convInt = 0
		} else {
			result = append(result, convInt)
		}
	}

	return result
}
