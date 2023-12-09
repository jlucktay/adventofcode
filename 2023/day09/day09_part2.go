// Package main for Advent of Code 2023, day 9, part 2
// https://adventofcode.com/2023/day/9
package main

func (h History) prevValue() int {
	if len(h) == 0 {
		return 0
	}

	deltas := h.reduceToZeroCount(true)

	yFactor := 0

	for i := len(deltas) - 2; i >= 0; i-- {
		xFactor := deltas[i]
		zFactor := xFactor - yFactor
		yFactor = zFactor
	}

	return h[0] - yFactor
}

func Part2(inputLines []string) (int, error) {
	puzzle, err := parseInput(inputLines)
	if err != nil {
		return 0, err
	}

	total := 0

	for i := 0; i < len(puzzle); i++ {
		next := puzzle[i].prevValue()

		total += next
	}

	return total, nil
}
