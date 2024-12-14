// Package day01 for Advent of Code 2019, day 1.
// https://adventofcode.com/2019/day/1
package day01

import (
	"fmt"
	"strconv"
	"strings"
)

type moduleMass int

type spacecraft []moduleMass

func parseInput(input string) (spacecraft, error) {
	result := make(spacecraft, 0)

	for _, line := range strings.Fields(input) {
		newModule, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("parsing '%s': %w", line, err)
		}

		result = append(result, moduleMass(newModule))
	}

	return result, nil
}
