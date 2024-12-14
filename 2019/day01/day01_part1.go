// Package day01 for Advent of Code 2019, day 1, part 1.
// https://adventofcode.com/2019/day/1
package day01

func Part1(input string) (int, error) {
	sc, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := 0

	for _, module := range sc {
		result += module.fuelRequired()
	}

	return result, nil
}

// Fuel required to launch a given module is based on its mass.
// Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.
func (mm moduleMass) fuelRequired() int {
	return (int(mm) / 3) - 2
}
