// Package day03 for Advent of Code 2024, day 3, part 2.
// https://adventofcode.com/2024/day/3
package day03

func Part2(input string) (int, error) {
	parsed, err := parseInput(input, true)
	if err != nil {
		return 0, err
	}

	result := 0

	for _, multi := range parsed {
		result += multi.left * multi.right
	}

	return result, nil
}
