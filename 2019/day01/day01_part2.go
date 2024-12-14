// Package day01 for Advent of Code 2019, day 1, part 2.
// https://adventofcode.com/2019/day/1
package day01

func Part2(input string) (int, error) {
	sc, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := 0

	for _, module := range sc {
		totalModuleFuel := 0

		for iteration := module.fuelRequired(); iteration > 0; iteration = moduleMass(iteration).fuelRequired() {
			totalModuleFuel += iteration
		}

		result += totalModuleFuel
	}

	return result, nil
}
