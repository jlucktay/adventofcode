// Package day13 for Advent of Code 2024, day 13, part 2.
// https://adventofcode.com/2024/day/13
package day13

func Part2(input string) (int, error) {
	machines, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := 0

	for _, machine := range machines {
		machine.p.X += 10_000_000_000_000
		machine.p.Y += 10_000_000_000_000

		result += calculate(machine)
	}

	return result, nil
}
