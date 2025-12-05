// Package day05 for Advent of Code 2025, day 5, part 1.
// https://adventofcode.com/2025/day/5
package day05

import "log/slog"

func Part1(input string) (int, error) {
	cafe, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	slog.Debug("cafe", slog.Any("raw", cafe))

	result := cafe.countFreshFromAvailable()

	return result, nil
}
