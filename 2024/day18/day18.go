// Package day18 for Advent of Code 2024, day 18.
// https://adventofcode.com/2024/day/18
package day18

import (
	"bufio"
	"bytes"
	"slices"
	"strings"
)

func parseInput(input string) (struct{}, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := struct{}{}

	for scanner.Scan() {
		xLine := strings.Split(scanner.Text(), " ")
		xLine = slices.DeleteFunc(xLine, func(s string) bool { return s == "" })

		// ...
		// go through input line by line and roll up into result
		// ...

    _ = xLine
	}

	// ...
	// validate parsed result
	// ...

	return result, nil
}
