// Package day08 for Advent of Code 2024, day 8, part 1.
// https://adventofcode.com/2024/day/8
package day08

import (
	"slices"
)

func Part1(input string) (int, error) {
	ag, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	ag.plotAntinodes()

	return len(ag.thingDirectory[Antinode]), nil
}

func (ag AntennaGrid) bounds() (int, int) {
	if len(ag.theGrid) > 0 && len(ag.theGrid[0]) > 0 {
		return len(ag.theGrid[0]), len(ag.theGrid)
	}

	return -1, -1
}

func (ag *AntennaGrid) plotAntinode(x, y int) {
	if x < 0 || y < 0 {
		return
	}

	xLimit, yLimit := ag.bounds()

	if x >= xLimit || y >= yLimit {
		return
	}

	// Add an antinode in the directory at this grid position.
	if _, inDirectory := ag.thingDirectory[Antinode]; !inDirectory {
		ag.thingDirectory[Antinode] = make([][2]int, 0)
	}

	plotHere := [2]int{x, y}

	if !slices.Contains(ag.thingDirectory[Antinode], plotHere) {
		ag.thingDirectory[Antinode] = append(ag.thingDirectory[Antinode], plotHere)
	}

	// If this grid position doesn't already have an antinode, add it.
	if slices.Contains(ag.theGrid[y][x], Antinode) {
		return
	}

	ag.theGrid[y][x] = append(ag.theGrid[y][x], Antinode)
}

func (ag *AntennaGrid) plotPairs(gpt GridPointThing, primary [2]int, secondaries [][2]int) {
	for _, secondary := range secondaries {
		xDiff, yDiff := secondary[0]-primary[0], secondary[1]-primary[1]

		ag.plotAntinode(primary[0]-xDiff, primary[1]-yDiff)
		ag.plotAntinode(secondary[0]+xDiff, secondary[1]+yDiff)
	}

	if len(secondaries) >= 2 {
		ag.plotPairs(gpt, secondaries[0], secondaries[1:])
	}
}

func (ag *AntennaGrid) plotAntinodes() {
	for gpt, gptLocations := range ag.thingDirectory {
		if gpt == Empty || gpt == Antinode {
			continue
		}

		if len(gptLocations) <= 1 {
			// Nothing to pair with.
			continue
		}

		ag.plotPairs(gpt, gptLocations[0], gptLocations[1:])
	}
}
