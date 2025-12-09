// Package day09 for Advent of Code 2025, day 9, part 1.
// https://adventofcode.com/2025/day/9
package day09

func Part1(input string) (int, error) {
	mt, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := mt.largestRectangle()

	return result, nil
}
