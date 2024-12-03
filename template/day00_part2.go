// Package day00 for Advent of Code <year>, day <day>, part 2.
// https://adventofcode.com/<year>/day/<day>
package day00

func Part2(input string) (int, error) {
	_, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := 0

	for range 27 {
		result += 42
	}

	return result, nil
}
