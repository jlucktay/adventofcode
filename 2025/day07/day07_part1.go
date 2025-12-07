// Package day07 for Advent of Code 2025, day 7, part 1.
// https://adventofcode.com/2025/day/7
package day07

func Part1(input string) (int, error) {
	tbs, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := tbs.startSplitting()

	return result, nil
}
