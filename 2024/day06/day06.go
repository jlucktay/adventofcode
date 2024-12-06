// Package day06 for Advent of Code 2024, day 6.
// https://adventofcode.com/2024/day/6
package day06

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
)

func parseInput(input string) (Map, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := Map{
		guardX:    -1,
		guardY:    -1,
		floorPlan: make([][]MapPoint, 0),
		visited:   make(map[[2]int]map[MapPoint]struct{}),
	}

	mapRow := 0
	guardInitialLocationSet := false

	for scanner.Scan() {
		result.floorPlan = append(result.floorPlan, make([]MapPoint, 0))

		for x, newPoint := range scanner.Text() {
			parsedPoint := MapPoints.Parse(newPoint)

			if parsedPoint == nil {
				return Map{}, fmt.Errorf("parsing '%v' as MapPoint", newPoint)
			}

			result.floorPlan[mapRow] = append(result.floorPlan[mapRow], *parsedPoint)

			if GuardMapPoints.Contains(*parsedPoint) {
				if guardInitialLocationSet {
					return Map{}, errors.New("attempt to set initial guard location more than once")
				}

				guardInitialLocationSet = true
				result.guardX = x
				result.guardY = mapRow
			}
		}

		mapRow++
	}

	return result, nil
}
