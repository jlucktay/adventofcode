// Package main for Advent of Code 2023, day 8, part 2
// https://adventofcode.com/2023/day/8
package main

type Ghost struct {
	current    Address
	stepsTaken uint64
}

type Ghosts []Ghost

func Part2(inputLines []string) (uint64, error) {
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

func gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(x ...uint64) uint64 {
	if len(x) == 0 {
		return 0
	} else if len(x) == 1 {
		return x[0]
	} else if len(x) == 2 {
		return x[0] * x[1] / gcd(x[0], x[1])
	}

	return lcm(x[0], lcm(x[1:]...))
}
