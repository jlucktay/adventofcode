// Package day02 for Advent of Code 2025, day 2.
// https://adventofcode.com/2025/day/2
package day02

import (
	"bufio"
	"bytes"
	"fmt"
	"log/slog"
	"slices"
	"strconv"
	"strings"
)

type ProductIDRange struct {
	start, finish int
}

func (pir *ProductIDRange) InvalidIDs() ([]int, error) {
	result := make([]int, 0)

	for i := pir.start; i <= pir.finish; i++ {
		s := strconv.Itoa(i)

		// Only keep going if the product ID contains an even number of digits.
		if len(s)%2 != 0 {
			continue
		}

		r := []rune(s)

		leftSide := r[:len(s)/2]
		rightSide := r[len(s)/2:]

		slog.Debug("split string", slog.String("left", string(leftSide)), slog.String("right", string(rightSide)))

		if strings.EqualFold(string(leftSide), string(rightSide)) {
			result = append(result, i)
		}
	}

	return result, nil
}

type GiftShopComputer struct {
	ranges []ProductIDRange
}

func parseInput(input string) (GiftShopComputer, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := GiftShopComputer{
		ranges: make([]ProductIDRange, 0),
	}

	for scanner.Scan() {
		xLine := strings.Split(scanner.Text(), " ")
		xLine = slices.DeleteFunc(xLine, func(s string) bool { return s == "" })

		if len(xLine) != 1 {
			return GiftShopComputer{}, fmt.Errorf("malformed line '%+v'", xLine)
		}

		productIDRanges := strings.Split(xLine[0], ",")

		for _, pir := range productIDRanges {
			xPIR := strings.Split(pir, "-")

			if len(xPIR) != 2 {
				return GiftShopComputer{}, fmt.Errorf("malformed product ID range '%s'", pir)
			}

			intStart, err := strconv.Atoi(xPIR[0])
			if err != nil {
				return GiftShopComputer{}, fmt.Errorf("converting start '%s' to int: %w", xPIR[0], err)
			}

			intFinish, err := strconv.Atoi(xPIR[1])
			if err != nil {
				return GiftShopComputer{}, fmt.Errorf("converting finish '%s' to int: %w", xPIR[1], err)
			}

			newPIR := ProductIDRange{start: intStart, finish: intFinish}

			result.ranges = append(result.ranges, newPIR)
		}
	}

	return result, nil
}
