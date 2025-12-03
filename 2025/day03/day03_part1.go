// Package day03 for Advent of Code 2025, day 3, part 1.
// https://adventofcode.com/2025/day/3
package day03

import "log/slog"

func Part1(input string) (int, error) {
	bb, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	slog.Debug("bb", slog.Any("bb", bb))

	result := 0

	for _, b := range bb {
		result += b.largestJoltage()
	}

	return result, nil
}
