// Package day01 for Advent of Code 2025, day 1, part 2.
// https://adventofcode.com/2025/day/1
package day01

func Part2(input string) (int, error) {
	rotations, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	safe := NewSafe()
	safe.Follow(rotations)

	return safe.anyClickPointsAtZero, nil
}
