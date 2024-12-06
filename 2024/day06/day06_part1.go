// Package day06 for Advent of Code 2024, day 6, part 1.
// https://adventofcode.com/2024/day/6
package day06

import (
	"strings"

	"github.com/orsinium-labs/enum"
)

func Part1(input string) (int, error) {
	myMap, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	myMap.FollowProtocol()

	result := 0

	return result, nil
}

type MapPoint enum.Member[rune]

var (
	Empty       = MapPoint{'.'}
	Obstruction = MapPoint{'#'}

	GuardNorth = MapPoint{'^'}
	GuardEast  = MapPoint{'>'}
	GuardSouth = MapPoint{'v'}
	GuardWest  = MapPoint{'<'}

	MapPoints      = enum.New(Empty, Obstruction, GuardNorth, GuardEast, GuardSouth, GuardWest)
	GuardMapPoints = enum.New(GuardNorth, GuardEast, GuardSouth, GuardWest)
)

type Map struct {
	guardX, guardY int

	floorPlan [][]MapPoint
}

func (m Map) String() string {
	output := strings.Builder{}

	for _, mapRow := range m.floorPlan {
		for _, mapPoint := range mapRow {
			_, _ = output.WriteRune(mapPoint.Value)
		}

		output.WriteString("\n")
	}

	return output.String()
}

func (m Map) FollowProtocol() {}
