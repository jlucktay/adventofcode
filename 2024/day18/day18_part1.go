// Package day18 for Advent of Code 2024, day 18, part 1.
// https://adventofcode.com/2024/day/18
package day18

import (
	"fmt"
	"image"
	"strconv"
)

func Part1(input string) (int, error) {
	rr, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	return rr.MinimumStepsToExit()
}

func (rr RAMRun) MinimumStepsToExit() (int, error) {
	start := image.Pt(0, 0)

	var finish image.Point

	switch rr.byteDropLimit {
	case -1:
		return 0, nil

	case 12:
		finish = image.Pt(6, 6)

	case 1024:
		finish = image.Pt(70, 70)

	default:
		panic("unknown byte drop limit: " + strconv.Itoa(rr.byteDropLimit))
	}

	path, err := rr.graph.Shortest(start, finish)
	if err != nil {
		return 0, fmt.Errorf("finding shortest path: %w", err)
	}

	return int(path.Distance), nil
}
