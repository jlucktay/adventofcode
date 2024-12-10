// Package day10 for Advent of Code 2024, day 10.
// https://adventofcode.com/2024/day/10
package day10

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"

	"github.com/RyanCarrier/dijkstra/v2"
)

type HikingMapPoint struct {
	height int
	name   string
}

// HikingMap is a thin wrapper over [dijkstra.MappedGraph] plus a raw representation of the hiking map to help with
// initial parsing.
type HikingMap struct {
	graph dijkstra.MappedGraph[string]
	raw   [][]HikingMapPoint

	trailheads, peaks []string
}

func parseInput(input string) (HikingMap, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := &HikingMap{
		raw:        make([][]HikingMapPoint, 0),
		graph:      dijkstra.NewMappedGraph[string](),
		trailheads: make([]string, 0),
		peaks:      make([]string, 0),
	}

	y := 0
	pointNames := make(map[rune]int)

	// First, get the raw height data into the 2-dimensional slice.
	for scanner.Scan() {
		line := scanner.Text()

		result.raw = append(result.raw, make([]HikingMapPoint, 0))

		for _, r := range line {
			var height int

			if r == '.' {
				height = -1
			} else {
				var err error
				height, err = strconv.Atoi(string(r))
				if err != nil {
					return HikingMap{}, fmt.Errorf("parsing '%s': %w", string(r), err)
				}
			}

			newPoint := HikingMapPoint{height: height}

			// Make up a string name for this particular point.
			newPoint.name = fmt.Sprintf("%dth-%s", pointNames[r], string(r))
			pointNames[r]++

			result.raw[y] = append(result.raw[y], newPoint)

			if err := result.graph.AddEmptyVertex(newPoint.name); err != nil {
				return HikingMap{}, fmt.Errorf("adding new vertex named '%s': %w", newPoint.name, err)
			}

			switch height {
			case 0:
				result.trailheads = append(result.trailheads, newPoint.name)

			case 9:
				result.peaks = append(result.peaks, newPoint.name)
			}
		}

		y++
	}

	// Now traverse the raw grid, look around from each point for valid neighbours, and populate the graph.
	for y := range result.raw {
		for x := range result.raw[y] {
			srcVertex := result.raw[y][x].name

			for _, dstVertex := range result.parseGetValidNeighbours(x, y) {
				if err := result.graph.AddArc(srcVertex, dstVertex, 1); err != nil {
					return HikingMap{}, fmt.Errorf("adding src '%s' and dst '%s': %w", srcVertex, dstVertex, err)
				}
			}
		}
	}

	return *result, nil
}

// > Based on un-scorched scraps of the book, you determine that a good hiking trail is as long as possible and has an
// > even, gradual, uphill slope. For all practical purposes, this means that a hiking trail is any path that starts at
// > height 0, ends at height 9, and always increases by a height of exactly 1 at each step. Hiking trails never
// > include diagonal steps - only up, down, left, or right (from the perspective of the map).
func (hm HikingMap) parseGetValidNeighbours(x, y int) []string {
	result := make([]string, 0)

	for _, yDelta := range []int{-1, 1} {
		if y+yDelta >= 0 && y+yDelta < len(hm.raw) && x >= 0 && x < len(hm.raw[y+yDelta]) {
			if hm.raw[y][x].height+1 == hm.raw[y+yDelta][x].height {
				result = append(result, hm.raw[y+yDelta][x].name)
			}
		}
	}

	for _, xDelta := range []int{-1, 1} {
		if y >= 0 && y < len(hm.raw) && x+xDelta >= 0 && x+xDelta < len(hm.raw[y]) {
			if hm.raw[y][x].height+1 == hm.raw[y][x+xDelta].height {
				result = append(result, hm.raw[y][x+xDelta].name)
			}
		}
	}

	return result
}
