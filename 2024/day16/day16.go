// Package day16 for Advent of Code 2024, day 16.
// https://adventofcode.com/2024/day/16
package day16

import (
	"fmt"
	"image"
	"strings"

	"github.com/orsinium-labs/enum"
)

type Tile enum.Member[rune]

var (
	Start = Tile{'S'}
	End   = Tile{'E'}
	Floor = Tile{'.'}
	Wall  = Tile{'#'}

	Tiles = enum.New(Start, End, Floor, Wall)
)

type ReindeerMaze struct {
	grid  map[image.Point]Tile
	start image.Point

	xMax, yMax int
}

func (rm ReindeerMaze) String() string {
	result := ""

	for y := range rm.yMax {
		for x := range rm.xMax {
			result += string(rm.grid[image.Pt(x, y)].Value)
		}

		result += "\n"
	}

	return result
}

func parseInput(input string) (ReindeerMaze, error) {
	if len(input) == 0 {
		return ReindeerMaze{}, nil
	}

	lines := strings.Fields(input)

	result := ReindeerMaze{
		grid: make(map[image.Point]Tile),
	}

	for y, line := range lines {
		for x, r := range line {
			parsed := Tiles.Parse(r)
			if parsed == nil {
				return ReindeerMaze{}, fmt.Errorf("unknown enum/rune '%s' at %d,%d", string(r), x, y)
			}

			here := image.Pt(x, y)

			if *parsed == Start {
				result.start = here
			}

			result.grid[here] = *parsed

			if result.xMax < x+1 {
				result.xMax = x + 1
			}
		}

		if result.yMax < y+1 {
			result.yMax = y + 1
		}
	}

	return result, nil
}
