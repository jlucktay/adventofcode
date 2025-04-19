// Package day09 for Advent of Code 2024, day 9.
// https://adventofcode.com/2024/day/9
package day09

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
)

func parseInput(input string) (Disk, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := make(Disk, 0)

	for scanner.Scan() {
		line := scanner.Text()

		for index, size := range line {
			parsedSize, err := strconv.ParseInt(string(size), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("parsing '%s': %w", string(size), err)
			}

			for range parsedSize {
				switch index%2 == 0 {
				case true:
					result = append(result, int64(index/2))

				case false:
					result = append(result, -1)
				}
			}
		}
	}

	return result, nil
}
