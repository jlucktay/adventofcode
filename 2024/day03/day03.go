// Package day03 for Advent of Code 2024, day 3.
// https://adventofcode.com/2024/day/3
package day03

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

type Multiplication struct {
	left, right int
}

type Multiplications []Multiplication

func parseInput(input string, part2 bool) (Multiplications, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	part1Pattern := `mul\(([0-9]{1,3}),([0-9]{1,3})\)`
	part2Pattern := `(mul\(([0-9]{1,3}),([0-9]{1,3})\)|do(n't)?\(\))`

	result := make(Multiplications, 0)

	if !part2 {
		re := regexp.MustCompile(part1Pattern)

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
	} else {
		re := regexp.MustCompile(part2Pattern)

		do := true

		for scanner.Scan() {
			scanned := re.FindAllStringSubmatch(scanner.Text(), -1)

			for _, reMatch := range scanned {
				if strings.HasPrefix(reMatch[0], "do()") {
					do = true
				} else if strings.HasPrefix(reMatch[0], "don't()") {
					do = false
				} else if strings.HasPrefix(reMatch[0], "mul(") {
					if !do {
						continue
					}

					l, err := strconv.Atoi(reMatch[2])
					if err != nil {
						return nil, err
					}

					r, err := strconv.Atoi(reMatch[3])
					if err != nil {
						return nil, err
					}

					result = append(result, Multiplication{left: l, right: r})
				}
			}
		}
	}

	return result, nil
}
