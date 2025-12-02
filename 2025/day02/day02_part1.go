// Package day02 for Advent of Code 2025, day 2, part 1.
// https://adventofcode.com/2025/day/2
package day02

import (
	"fmt"
	"log/slog"
)

func Part1(input string) (int, error) {
	gsc, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := 0

	for _, gr := range gsc.ranges {
		slog.Debug("productIDRange", slog.Any("gr", gr))

		grii, err := gr.InvalidIDs()
		if err != nil {
			return 0, fmt.Errorf("calculating invalid IDs: %w", err)
		}

		for _, invalidID := range grii {
			slog.Debug("invalidIDs", slog.Any("invalidID", invalidID))

			result += invalidID
		}
	}

	return result, nil
}
