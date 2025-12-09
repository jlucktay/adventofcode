// Package day08 for Advent of Code 2025, day 8, part 1.
// https://adventofcode.com/2025/day/8
package day08

import (
	"log/slog"
)

func Part1(input string) (int, error) {
	pg, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	slog.Debug("calculating distances between boxes", slog.Int("boxes", len(pg.boxes)))
	pg.calculateDistances()
	slog.Debug("calculated distances between boxes", slog.Int("boxes", len(pg.boxes)), slog.Int("distancePairs", len(pg.distances)))

	slog.Debug("connecting closest pairs into circuits", slog.Int("boxes", len(pg.boxes)))
	pg.connectClosestPairs()
	slog.Debug("connected closest pairs into circuits", slog.Int("boxes", len(pg.boxes)), slog.Int("circuits", len(pg.circuits)))

	result := pg.largestThreeCircuits()

	return result, nil
}
