// Package day05 for Advent of Code 2024, day 5.
// https://adventofcode.com/2024/day/5
package day05

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) (manualPrinter, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := manualPrinter{
		rules:   make([]rule, 0),
		updates: make([]update, 0),
	}

	parsingUpdates := false

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			parsingUpdates = true

			continue
		}

		if !parsingUpdates {
			xLine := strings.Split(line, "|")

			if len(xLine) != 2 {
				return manualPrinter{}, errors.New("rule does not have 2 components")
			}

			left, err := strconv.Atoi(xLine[0])
			if err != nil {
				return manualPrinter{}, fmt.Errorf("could not parse '%s' into an integer: %w", xLine[0], err)
			}

			right, err := strconv.Atoi(xLine[1])
			if err != nil {
				return manualPrinter{}, fmt.Errorf("could not parse '%s' into an integer: %w", xLine[1], err)
			}

			newRule := [2]int{left, right}
			result.rules = append(result.rules, newRule)
		}

		if parsingUpdates {
			xLine := strings.Split(line, ",")

			newUpdate := make(update, 0)

			for _, pageNumber := range xLine {
				pageNum, err := strconv.Atoi(pageNumber)
				if err != nil {
					return manualPrinter{}, fmt.Errorf("could not parse '%s' into an integer: %w", pageNumber, err)
				}

				newUpdate = append(newUpdate, pageNum)
			}

			result.updates = append(result.updates, newUpdate)
		}
	}

	return result, nil
}
