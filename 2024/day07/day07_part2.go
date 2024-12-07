// Package day07 for Advent of Code 2024, day 7, part 2.
// https://adventofcode.com/2024/day/7
package day07

func Part2(input string) (uint64, error) {
	xEq, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := uint64(0)

	for _, eq := range xEq {
		if attemptResult, found, err := attempt(eq.desiredResult, eq.numbers[1:], eq.numbers[0], true); err == nil &&
			found && attemptResult > 0 {

			result += eq.desiredResult
		}
	}

	return result, nil
}

var Concat = Operator{'|'}
