// Package day05 for Advent of Code 2025, day 5.
// https://adventofcode.com/2025/day/5
package day05

import (
	"bufio"
	"bytes"
	"cmp"
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

func (c Cafeteria) countAllFresh() int {
	slog.Debug("start counting all fresh ingredients")

	result := 0

	defer func() { slog.Debug("finish counting all fresh ingredients", slog.Int("result", result)) }()

	if len(c.fresh) == 0 {
		return 0
	}

	// 1. Sort all number ranges based on their starting points
	slices.SortStableFunc(c.fresh, func(a, b freshRange) int {
		if a.lowerBound < b.lowerBound {
			return -1
		}

		if a.lowerBound > b.lowerBound {
			return 1
		}

		return cmp.Compare(a.upperBound, b.upperBound)
	})

	slog.Debug("fresh ranges, sorted", slog.String("slice", fmt.Sprintf("%+v", c.fresh)))

	// 2. Create an empty list for merged ranges and add the first sorted range to it
	mergedRanges := []freshRange{c.fresh[0]}

	// 3. Iterate and merge
	for i := 1; i < len(c.fresh); i++ {
		// 1. Take the next range from the sorted list
		nextRange := c.fresh[i]

		slog.Debug("comparing two fresh ranges", slog.String("lastMerged", fmt.Sprintf("%+v", mergedRanges[len(mergedRanges)-1])), slog.String("next", fmt.Sprintf("%+v", nextRange)))

		// 2. Compare it with the last range in the merged ranges
		if mergedRanges[len(mergedRanges)-1].overlapsWith(nextRange) {
			// Update the end of the last merged range to be the maximum of both ends (effectively merging them)
			if nextRange.lowerBound < mergedRanges[len(mergedRanges)-1].lowerBound {
				mergedRanges[len(mergedRanges)-1].lowerBound = nextRange.lowerBound
			}

			if nextRange.upperBound > mergedRanges[len(mergedRanges)-1].upperBound {
				mergedRanges[len(mergedRanges)-1].upperBound = nextRange.upperBound
			}

			slog.Debug("updated last merged", slog.String("lastMerged", fmt.Sprintf("%+v", mergedRanges[len(mergedRanges)-1])))
		} else {
			// Add the current range as a new, separate range to the merged ranges
			mergedRanges = append(mergedRanges, nextRange)

			slog.Debug("added new merged", slog.String("next", fmt.Sprintf("%+v", nextRange)))
		}
	}

	// 4. Calculate total coverage
	slog.Debug("calculating total coverage", slog.String("mergedRanges", fmt.Sprintf("%+v", mergedRanges)))

	for mrIndex := range mergedRanges {
		singleRangeCoverage := mergedRanges[mrIndex].upperBound - mergedRanges[mrIndex].lowerBound + 1

		slog.Debug("calculated coverage from single range", slog.String("mergedRanges[mrIndex]", fmt.Sprintf("%+v", mergedRanges[mrIndex])), slog.Int("coverage", singleRangeCoverage))

		result += singleRangeCoverage
	}

	return result
}

func (fr freshRange) overlapsWith(target freshRange) bool {
	if fr.upperBound < target.lowerBound || target.upperBound < fr.lowerBound {
		return false
	}

	return target.lowerBound <= fr.upperBound
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
