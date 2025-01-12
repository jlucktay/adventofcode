// Package day17 for Advent of Code 2024, day 17, part 1.
// https://adventofcode.com/2024/day/17
package day17

func Part1(input string) (string, error) {
	cc, err := parseInput(input)
	if err != nil {
		return "", err
	}

	for cc.Process() {
		// Run until halt.
	}

	return cc.Output(), nil
}
