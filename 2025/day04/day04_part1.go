// Package day04 for Advent of Code 2025, day 4, part 1.
// https://adventofcode.com/2025/day/4
package day04

import (
	"fmt"
)

func Part1(input string) (int, error) {
	pd, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	fmt.Printf("%s\n", pd)

	return pd.paperAccessibleByForklift(), nil
}
