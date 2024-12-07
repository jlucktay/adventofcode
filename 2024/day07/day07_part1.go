// Package day07 for Advent of Code 2024, day 7, part 1.
// https://adventofcode.com/2024/day/7
package day07

import "github.com/orsinium-labs/enum"

func Part1(input string) (uint64, error) {
	xEq, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := uint64(0)

	for _, eq := range xEq {
		if attemptResult, found := attempt(eq.desiredResult, eq.numbers[1:], eq.numbers[0]); found && attemptResult > 0 {
			result += eq.desiredResult
		}
	}

	return result, nil
}

type Operator enum.Member[rune]

var (
	Add  = Operator{'+'}
	Mult = Operator{'*'}

	Operators = enum.New(Add, Mult)
)

func attempt(desired uint64, inputs []uint64, subtotal uint64) (uint64, bool) {
	if subtotal > desired {
		return 0, false
	}

	for _, op := range Operators.Members() {
		iterTotal := subtotal

		switch op {
		case Add:
			iterTotal += inputs[0]

		case Mult:
			iterTotal *= inputs[0]

		default:
			return 0, false
		}

		if iterTotal == desired && len(inputs) == 1 {
			return iterTotal, true
		}

		if len(inputs) > 1 {
			if result, found := attempt(desired, inputs[1:], iterTotal); found {
				return result, true
			}
		}
	}

	return 0, false
}
