// Package day05 for Advent of Code 2025, day 5.
// https://adventofcode.com/2025/day/5
package day05

import (
	"bufio"
	"bytes"
	"fmt"
	"log/slog"
	"slices"
	"strconv"
	"strings"
)

type Cafeteria struct {
	fresh     []freshRange
	available []int
}

type freshRange struct {
	lowerBound, upperBound int
}

func (c Cafeteria) countFreshFromAvailable() int {
	result := 0

	slog.Debug("counting", slog.Int("available", len(c.available)))

	for availIndex := range c.available {
		slog.Debug("checking", slog.Int("available", c.available[availIndex]))

		for freshIndex := range c.fresh {
			if c.fresh[freshIndex].contains(c.available[availIndex]) {
				result++

				break
			}
		}
	}

	return result
}

func (fr freshRange) contains(target int) bool {
	return target >= fr.lowerBound && target <= fr.upperBound
}

func parseInput(input string) (Cafeteria, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := Cafeteria{
		fresh:     make([]freshRange, 0),
		available: make([]int, 0),
	}

	for scanner.Scan() {
		st := scanner.Text()

		slog.Debug("scanner", slog.String("text", st))

		xLine := strings.Split(st, "-")
		xLine = slices.DeleteFunc(xLine, func(s string) bool { return s == "" })

		slog.Debug("xLine", slog.Int("len", len(xLine)), slog.String("raw", fmt.Sprintf("%+v", xLine)))

		switch len(xLine) {
		case 0:
			slog.Debug("case 0")

			continue

		case 1:
			slog.Debug("case 1", slog.String("line", fmt.Sprintf("%+v", xLine)))

			// Available ingredient IDs
			availableID, err := strconv.Atoi(xLine[0])
			if err != nil {
				return Cafeteria{}, fmt.Errorf("converting available ingredient ID '%s': %w", xLine[0], err)
			}

			result.available = append(result.available, availableID)

		case 2:
			slog.Debug("case 2", slog.String("line", fmt.Sprintf("%+v", xLine)))

			// Fresh ingredient ID ranges
			lowerBound, err := strconv.Atoi(xLine[0])
			if err != nil {
				return Cafeteria{}, fmt.Errorf("converting lower bound '%s': %w", xLine[0], err)
			}

			upperBound, err := strconv.Atoi(xLine[1])
			if err != nil {
				return Cafeteria{}, fmt.Errorf("converting upper bound '%s': %w", xLine[1], err)
			}

			slog.Debug("parsed lower and upper bounds", slog.Int("size", upperBound-lowerBound+1), slog.Int("lower", lowerBound), slog.Int("upper", upperBound))

			newFreshRange := freshRange{lowerBound: lowerBound, upperBound: upperBound}
			result.fresh = append(result.fresh, newFreshRange)

			slog.Debug("stored range in fresh slice", slog.Int("lower", lowerBound), slog.Int("upper", upperBound))

		default:
			slog.Debug("default")

			return Cafeteria{}, fmt.Errorf("unexpected line format '%+v'", xLine)
		}
	}

	slices.Sort(result.available)
	result.available = slices.Compact(result.available)

	return result, nil
}
