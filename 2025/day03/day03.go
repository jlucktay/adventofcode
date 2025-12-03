// Package day03 for Advent of Code 2025, day 3.
// https://adventofcode.com/2025/day/3
package day03

import (
	"bufio"
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type BatteryBank []int

type BatteryBanks []BatteryBank

func (bb BatteryBank) largestJoltage() int {
	highestTotal := 0

	for i := 0; i < len(bb)-1; i++ {
		left := bb[i]

		for j := i + 1; j < len(bb); j++ {
			right := bb[j]

			if candidate := left*10 + right; candidate > highestTotal {
				highestTotal = candidate
			}
		}
	}

	return highestTotal
}

func parseInput(input string) (BatteryBanks, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := BatteryBanks{}

	for scanner.Scan() {
		xLine := strings.Split(scanner.Text(), " ")
		xLine = slices.DeleteFunc(xLine, func(s string) bool { return s == "" })

		if len(xLine) != 1 {
			return nil, fmt.Errorf("unexpected format for line '%+v'", xLine)
		}

		newBank := make(BatteryBank, 0)

		for _, r := range xLine[0] {
			newBattery, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, fmt.Errorf("converting battery value '%v'", r)
			}

			newBank = append(newBank, newBattery)
		}

		result = append(result, newBank)
	}

	return result, nil
}
