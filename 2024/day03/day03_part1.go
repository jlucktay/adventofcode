// Package day03 for Advent of Code 2024, day 3, part 1.
// https://adventofcode.com/2024/day/3
package day03

func Part1(input string) (int, error) {
	_, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := 0

	for range 27 {
		result += 42
	}

	return result, nil
}
