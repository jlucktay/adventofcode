// Package day18 for Advent of Code 2024, day 18.
// https://adventofcode.com/2024/day/18
package day18

import (
	"fmt"
	"image"
	"strings"

	"github.com/RyanCarrier/dijkstra/v2"
	"github.com/orsinium-labs/enum"
)

type Tile enum.Member[rune]

var (
	Empty = Tile{'.'}
	Byte  = Tile{'#'}

	Tiles = enum.New(Empty, Byte)
)

type RAMRun struct {
	grid          map[image.Point]Tile
	bounds        image.Rectangle
	byteDropLimit int
	graph         dijkstra.MappedGraph[image.Point]

	undroppedRawBytes []string
}

func (rr RAMRun) String() string {
	result := ""

	for y := range rr.bounds.Max.Y {
		for x := range rr.bounds.Max.X {
			result += string(rr.grid[image.Pt(x, y)].Value)
		}

		result += "\n"
	}

	return result
}

func parseInput(input string) (RAMRun, error) {
	if len(input) == 0 {
		return RAMRun{byteDropLimit: -1}, nil
	}

	result := RAMRun{
		grid:  make(map[image.Point]Tile),
		graph: dijkstra.NewMappedGraph[image.Point](),
	}

	xLines := strings.Split(strings.TrimSpace(input), "\n")

	// Slightly different parameters in test and live modes.
	if len(xLines) <= 25 {
		result.bounds = image.Rect(0, 0, 7, 7)
		result.byteDropLimit = 12
	} else {
		result.bounds = image.Rect(0, 0, 71, 71)
		result.byteDropLimit = 1024
	}

	// Populate with empties to get started.
	for x := range result.bounds.Max.X {
		for y := range result.bounds.Max.Y {
			result.grid[image.Pt(x, y)] = Empty

			if err := result.graph.AddEmptyVertex(image.Pt(x, y)); err != nil {
				return RAMRun{}, fmt.Errorf("adding empty vertex at %d,%d to graph: %w", x, y, err)
			}
		}
	}

	// Start parsing coordinates of bytes.
	for i, line := range xLines {
		var x, y int

		n, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil || n != 2 {
			return RAMRun{}, fmt.Errorf("scanning line '%s' parsed %d items: %w", line, n, err)
		}

		newTile := image.Pt(x, y)
		result.grid[newTile] = Byte

		if i+1 >= result.byteDropLimit {
			result.undroppedRawBytes = xLines[i+1:]

			break
		}
	}

	// Look through the parsed grid for valid neighbours to feed into the graph.
	for x := range result.bounds.Max.X {
		for y := range result.bounds.Max.Y {
			position := image.Pt(x, y)

			for _, neighbour := range parseGetValidNeighbours(result.grid, position) {
				if err := result.graph.AddArc(position, neighbour, 1); err != nil {
					return RAMRun{}, fmt.Errorf("adding arc between %s and %s: %w", position, neighbour, err)
				}
			}
		}
	}

	return result, nil
}

func parseGetValidNeighbours(grid map[image.Point]Tile, position image.Point) []image.Point {
	result := make([]image.Point, 0)

	for _, direction := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
		candidate := position.Add(direction)

		if tileType, ok := grid[candidate]; ok && tileType == Empty {
			result = append(result, candidate)
		}
	}

	return result
}
