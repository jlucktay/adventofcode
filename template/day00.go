// Package day00 for Advent of Code <year>, day <day>.
// https://adventofcode.com/<year>/day/<day>
package day00

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
	}

	// ...
	// validate parsed result
	// ...

	return result, nil
}
