// Package main for Advent of Code 2023, day 8, part 2
// https://adventofcode.com/2023/day/8
package main

type Ghost struct {
	current    Address
	stepsTaken int
}

type Ghosts []Ghost

func Part2(inputLines []string) (int, error) {
	puzzle, err := parseLines(inputLines)
	if err != nil {
		return 0, err
	}

	result, err := puzzle.getToZZZ(true)
	if err != nil {
		return 0, err
	}

	return result, nil
}
