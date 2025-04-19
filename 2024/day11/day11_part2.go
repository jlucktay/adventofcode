// Package day11 for Advent of Code 2024, day 11, part 2.
// https://adventofcode.com/2024/day/11
package day11

func Part2(input string) (int, error) {
	stones, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	for range 75 {
		stones = stones.blink()
	}

	result := 0

	for _, count := range stones {
		result += count
	}

	return result, nil
}
