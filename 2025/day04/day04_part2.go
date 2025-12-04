// Package day04 for Advent of Code 2025, day 4, part 2.
// https://adventofcode.com/2025/day/4
package day04

import (
	"fmt"
	"log/slog"
)

func Part2(input string) (int, error) {
	pd, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	fmt.Printf("Initial state:\n%s\n", pd)

	runningTotal := 0

	for {
		removedThisTime := pd.paperAccessibleByForklift(true)

		runningTotal += removedThisTime

		slog.Debug("removed paper", slog.Int("removedThisTime", removedThisTime), slog.Int("runningTotal", runningTotal))

		fmt.Printf("Removed %d rolls of paper:\n%s\n", removedThisTime, pd)

		if removedThisTime == 0 {
			break
		}
	}

	return runningTotal, nil
}
