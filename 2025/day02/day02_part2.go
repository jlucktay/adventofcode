// Package day02 for Advent of Code 2025, day 2, part 2.
// https://adventofcode.com/2025/day/2
package day02

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

func (pir *ProductIDRange) InvalidIDsPart2() ([]int, error) {
	result := make([]int, 0)

	for i := pir.start; i <= pir.finish; i++ {
		s := strconv.Itoa(i)

		doubled := s + s

		if strings.Contains(doubled[1:len(doubled)-1], s) {
			result = append(result, i)
		}
	}

	return result, nil
}

func Part2(input string) (int, error) {
	gsc, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := 0

	for _, gr := range gsc.ranges {
		slog.Debug("productIDRange", slog.Any("gr", gr))

		grii, err := gr.InvalidIDsPart2()
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
