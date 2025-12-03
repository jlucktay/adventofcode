// Package day03 for Advent of Code 2025, day 3, part 2.
// https://adventofcode.com/2025/day/3
package day03

import (
	"fmt"
	"log/slog"
)

func Part2(input string) (int64, error) {
	bb, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	slog.Debug("bb", slog.Any("bb", bb))

	var result int64 = 0

	for _, b := range bb {
		osf, err := b.overcomeStaticFriction()
		if err != nil {
			return 0, fmt.Errorf("overcoming static friction: %w", err)
		}

		result += osf
	}

	return result, nil
}
