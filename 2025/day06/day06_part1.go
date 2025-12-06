// Package day06 for Advent of Code 2025, day 6, part 1.
// https://adventofcode.com/2025/day/6
package day06

func Part1(input string) (int, error) {
	tc, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := tc.grandTotal()

	return result, nil
}
