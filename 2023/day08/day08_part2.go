// Package main for Advent of Code 2023, day 8, part 2
// https://adventofcode.com/2023/day/8
package main

import (
	"slices"
)

type Ghost struct {
	current    Address
	stepsTaken uint64
}

type Ghosts []Ghost

func Part2(inputLines []string) (uint64, error) {
	puzzle, err := parseLines(inputLines)
	if err != nil {
		return 0, err
	}

	result, err := puzzle.getToZZZ(true)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func lcm(ghosts Ghosts) uint64 {
	originalValues := make([]uint64, 0)
	lcmValues := make([]uint64, 0)

	for _, ghost := range ghosts {
		originalValues = append(originalValues, ghost.stepsTaken)
		lcmValues = append(lcmValues, ghost.stepsTaken)
	}

	if len(lcmValues) == 0 {
		return 0
	}

	for {
		compactFodder := make([]uint64, len(lcmValues))

		if length := copy(compactFodder, lcmValues); length != len(lcmValues) {
			return 0
		}

		if compacted := slices.Compact(compactFodder); len(compacted) == 1 {
			return compacted[0]
		}

		lowestLCMValue := slices.Min(lcmValues)
		indexOfLowest := slices.Index(lcmValues, lowestLCMValue)

		originalValueToAdd := originalValues[indexOfLowest]
		newValueToSet := lcmValues[indexOfLowest] + originalValueToAdd

		lcmValues[indexOfLowest] = newValueToSet
	}
}
