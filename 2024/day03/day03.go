// Package day03 for Advent of Code 2024, day 3.
// https://adventofcode.com/2024/day/3
package day03

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
)

type Multiplication struct {
	left, right int
}

type Multiplications []Multiplication

func parseInput(input string) (Multiplications, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	result := make(Multiplications, 0)

	for scanner.Scan() {
		scanned := re.FindAllStringSubmatch(scanner.Text(), -1)

		for _, reMatch := range scanned {
			l, err := strconv.Atoi(reMatch[1])
			if err != nil {
				return nil, err
			}

			r, err := strconv.Atoi(reMatch[2])
			if err != nil {
				return nil, err
			}

			result = append(result, Multiplication{left: l, right: r})
		}
	}

	return result, nil
}
