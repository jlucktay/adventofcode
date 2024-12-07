// Package day07 for Advent of Code 2024, day 7, part 1.
// https://adventofcode.com/2024/day/7
package day07

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/orsinium-labs/enum"
)

func Part1(input string) (uint64, error) {
	xEq, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := uint64(0)

	for _, eq := range xEq {
		if attemptResult, found, err := attempt(eq.desiredResult, eq.numbers[1:], eq.numbers[0], false); err == nil &&
			found && attemptResult > 0 {

			result += eq.desiredResult
		}
	}

	return result, nil
}

type Operator enum.Member[rune]

var (
	Add  = Operator{'+'}
	Mult = Operator{'*'}

	Operators = enum.New(Add, Mult, Concat)
)

var (
	ErrSubtotalExceedDesired = errors.New("subtotal exceeds desired")
	ErrUnknownOperator       = errors.New("unknown operator")
	ErrExhausted             = errors.New("exhausted known operators")
)

func attempt(desired uint64, inputs []uint64, subtotal uint64, part2 bool) (uint64, bool, error) {
	if subtotal > desired {
		return 0, false, fmt.Errorf("%w: %d > %d", ErrSubtotalExceedDesired, subtotal, desired)
	}

	for _, op := range Operators.Members() {
		if !part2 && op == Concat {
			continue
		}

		iterTotal := subtotal

		switch op {
		case Add:
			iterTotal += inputs[0]

		case Mult:
			iterTotal *= inputs[0]

		case Concat:
			strCombined := fmt.Sprintf("%d%d", iterTotal, inputs[0])

			combinedNumber, err := strconv.ParseUint(strCombined, 10, 64)
			if err != nil {
				return 0, false, fmt.Errorf("parsing new number from '%s': %w", strCombined, err)
			}

			iterTotal = combinedNumber

		default:
			return 0, false, fmt.Errorf("%w: %s", ErrUnknownOperator, string(op.Value))
		}

		if iterTotal == desired && len(inputs) == 1 {
			return iterTotal, true, nil
		}

		if len(inputs) > 1 {
			if result, found, err := attempt(desired, inputs[1:], iterTotal, part2); found {
				return result, true, err
			}
		}
	}

	return 0, false, fmt.Errorf("%w: %v", ErrExhausted, Operators.Values())
}
