// Package day16 for Advent of Code 2024, day 16, part 2.
// https://adventofcode.com/2024/day/16
package day16

func Part2(input string) (int, error) {
	rm, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	if len(rm.grid) == 0 {
		return 0, nil
	}

	return rm.lowestScore(true), nil
}
