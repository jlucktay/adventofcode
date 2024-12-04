// Package day04 for Advent of Code 2024, day 4.
// https://adventofcode.com/2024/day/4
package day04

import (
	"bufio"
	"bytes"
)

type wordSearch [][]rune

func parseInput(input string) (wordSearch, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := make(wordSearch, 0)
	row := 0

	for scanner.Scan() {
		result = append(result, make([]rune, 0))

		for _, letter := range scanner.Text() {
			result[row] = append(result[row], letter)
		}

		row++
	}

	return result, nil
}
