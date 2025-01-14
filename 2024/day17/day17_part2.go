// Package day17 for Advent of Code 2024, day 17, part 2.
// https://adventofcode.com/2024/day/17
package day17

import (
	"slices"
)

func runToHalt(input string, a int) []int {
	cc, _ := parseInput(input)
	cc.A = a

	for cc.Process() {
	}

	return cc.rawOutput
}

func Part2(input string) (int, error) {
	cc, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	if len(cc.Program) == 0 {
		return 0, nil
	}

	initA := 0

	for n := len(cc.Program) - 1; n >= 0; n-- {
		initA <<= 3

		for !slices.Equal(runToHalt(input, initA), cc.Program[n:]) {
			initA++
		}
	}

	return initA, nil
}
