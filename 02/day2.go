package adventDay2

import (
	"strconv"
	"strings"
)

/*
ConvertInput will parse and rearrange the input provided for the day 2 exercises of the 2017 Advent of Code:
http://adventofcode.com/2017/day/2
*/
func ConvertInput(input string) [][]int {
	result := make([][]int, 16)

	for row, i := range strings.Split(input, "\n") {
		result[row] = make([]int, 16)

		for column, j := range strings.Split(i, "\t") {
			convInt, convErr := strconv.Atoi(strings.TrimSpace(string(j)))

			if convErr != nil {
				panic(convErr)
			}

			result[row][column] = convInt
		}
	}

	return result
}
