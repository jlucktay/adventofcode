// Package day05 for Advent of Code 2025, day 5, part 2.
// https://adventofcode.com/2025/day/5
package day05

func Part2(input string) (int, error) {
	cafe, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := cafe.countAllFresh()

	return result, nil
}
