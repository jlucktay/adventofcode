// Package day2 for Advent of Code 2017, day 2
// http://adventofcode.com/2017/day/2
// Contains a function used by both parts of day 2
package day2

import (
	"strconv"
	"strings"
)

/*
ConvertInput will parse and rearrange the input provided for the day 2 exercises of the 2017 Advent of Code:
http://adventofcode.com/2017/day/2
*/
func ConvertInput(input string) [][]int {
	result := make([][]int, 0)

	for row, i := range strings.Split(input, "\n") {
		result = append(result, make([]int, 0))

		for _, j := range strings.Split(i, "\t") {
			convInt, err := strconv.Atoi(strings.TrimSpace(string(j)))

			if err != nil {
				convInt = 0
			}

			result[row] = append(result[row], convInt)
		}
	}

	return result
}
