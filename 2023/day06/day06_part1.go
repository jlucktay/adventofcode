// Package main for Advent of Code 2023, day 6, part 1
// https://adventofcode.com/2023/day/6
package main

import (
	"slices"
	"strconv"
	"strings"
)

func parseInput(inputLines []string) []Race {
	times, distances := []int{}, []int{}

	for ilIndex := range inputLines {
		if trimmedTime, found := strings.CutPrefix(inputLines[ilIndex], "Time:"); found {
			xTimes := strings.Split(strings.TrimSpace(trimmedTime), " ")
			xTimes = slices.Compact(xTimes)

			for _, xt := range xTimes {
				if convertedTime, err := strconv.Atoi(xt); err == nil {
					times = append(times, convertedTime)
				}
			}
		}

		if trimmedDistance, found := strings.CutPrefix(inputLines[ilIndex], "Distance:"); found {
			xDistances := strings.Split(strings.TrimSpace(trimmedDistance), " ")
			xDistances = slices.Compact(xDistances)

			for _, xd := range xDistances {
				if convertedDistance, err := strconv.Atoi(xd); err == nil {
					distances = append(distances, convertedDistance)
				}
			}
		}
	}

	races := []Race{}

	for index := 0; index < len(times); index++ {
		races = append(races,
			Race{
				time:     times[index],
				distance: distances[index],
			})
	}

	return races
}

func Part1(inputLines []string) (int, error) {
	races := parseInput(inputLines)

	runningTotal := 0

	for _, race := range races {
		wtw := race.WaysToWin()
		if wtw > 0 {
			if runningTotal == 0 {
				runningTotal = wtw
			} else {
				runningTotal *= wtw
			}
		}
	}

	return runningTotal, nil
}

type Race struct {
	time, distance int
}

func (r Race) WaysToWin() int {
	timeLeftToTravel := r.time
	speed := 0
	result := 0

	for timeLeftToTravel > 0 {
		speed++
		timeLeftToTravel--

		distanceTraveled := speed * timeLeftToTravel

		if distanceTraveled > r.distance {
			result++
		}
	}

	return result
}
