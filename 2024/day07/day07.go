// Package day07 for Advent of Code 2024, day 7.
// https://adventofcode.com/2024/day/7
package day07

import (
	"bufio"
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type equation struct {
	desiredResult uint64
	numbers       []uint64
}

type equations []equation

func parseInput(input string) (equations, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := make(equations, 0)
	lineCounter := 0

	for scanner.Scan() {
		result = append(result, equation{numbers: make([]uint64, 0)})

		line := scanner.Text()
		colonIdx := strings.Index(line, ":")
		desiredResult, err := strconv.ParseUint(line[:colonIdx], 10, 64)
		if err != nil {
			return equations{}, fmt.Errorf("parsing desired result from '%s': %w", line[:colonIdx], err)
		}

		result[lineCounter].desiredResult = desiredResult

		xRawNumbers := strings.Split(line[colonIdx+1:], " ")
		xRawNumbers = slices.DeleteFunc(xRawNumbers, func(s string) bool { return s == "" })

		for _, rawNum := range xRawNumbers {
			newNumber, err := strconv.ParseUint(rawNum, 10, 64)
			if err != nil {
				return equations{}, fmt.Errorf("parsing new number from '%s': %w", rawNum, err)
			}

			result[lineCounter].numbers = append(result[lineCounter].numbers, newNumber)
		}

		lineCounter++
	}

	return result, nil
}
