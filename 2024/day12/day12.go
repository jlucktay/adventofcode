// Package day12 for Advent of Code 2024, day 12.
// https://adventofcode.com/2024/day/12
package day12

import (
	"bufio"
	"bytes"
	"image"
)

type Garden map[image.Point]rune

func parseInput(input string) (Garden, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := make(Garden)

	y := 0

	for scanner.Scan() {
		for x, pt := range scanner.Text() {
			result[image.Point{x, y}] = pt
		}

		y++
	}

	return result, nil
}
