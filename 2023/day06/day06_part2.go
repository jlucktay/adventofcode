// Package main for Advent of Code 2023, day 6, part 2
// https://adventofcode.com/2023/day/6
package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseInputPart2(inputLines []string) (Race, error) {
	times, distances := []string{}, []string{}

	for ilIndex := range inputLines {
		if trimmedTime, found := strings.CutPrefix(inputLines[ilIndex], "Time:"); found {
			xTimes := strings.Split(strings.TrimSpace(trimmedTime), " ")
			xTimes = slices.Compact(xTimes)

			for _, xt := range xTimes {
				if len(xt) > 0 {
					times = append(times, xt)
				}
			}
		}

		if trimmedDistance, found := strings.CutPrefix(inputLines[ilIndex], "Distance:"); found {
			xDistances := strings.Split(strings.TrimSpace(trimmedDistance), " ")
			xDistances = slices.Compact(xDistances)

			for _, xd := range xDistances {
				if len(xd) > 0 {
					distances = append(distances, xd)
				}
			}
		}
	}

	allTime := strings.Builder{}
	allDistance := strings.Builder{}

	for index := 0; index < len(times); index++ {
		allTime.WriteString(times[index])
	}

	for index := 0; index < len(distances); index++ {
		allDistance.WriteString(distances[index])
	}

	t, err := strconv.Atoi(allTime.String())
	if err != nil {
		return Race{}, fmt.Errorf("converting '%s': %w", allTime.String(), err)
	}

	d, err := strconv.Atoi(allDistance.String())
	if err != nil {
		return Race{}, fmt.Errorf("converting '%s': %w", allDistance.String(), err)
	}

	return Race{time: t, distance: d}, nil
}

func Part2(inputLines []string) (int, error) {
	race, err := parseInputPart2(inputLines)
	if err != nil {
		return 0, err
	}

	return race.WaysToWin(), nil
}
