// Package day11 for Advent of Code 2024, day 11.
// https://adventofcode.com/2024/day/11
package day11

import (
	"bufio"
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Stone uint64

type Stones map[Stone]int

func parseInput(input string) (Stones, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := make(Stones, 0)

	for scanner.Scan() {
		xLine := strings.Split(scanner.Text(), " ")
		xLine = slices.DeleteFunc(xLine, func(s string) bool { return s == "" })

		for _, stone := range xLine {
			newStone, err := strconv.ParseUint(stone, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("converting stone '%s': %w", stone, err)
			}

			result[Stone(newStone)]++
		}
	}

	return result, nil
}
