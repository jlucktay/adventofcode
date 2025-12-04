// Package day04 for Advent of Code 2025, day 4.
// https://adventofcode.com/2025/day/4
package day04

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"log/slog"
	"slices"
	"strings"

	"github.com/orsinium-labs/enum"
)

type Tile enum.Member[rune]

var (
	Empty = Tile{'.'}
	Paper = Tile{'@'}

	Tiles = enum.New(Empty, Paper)
)

type PrintingDepartment struct {
	grid       map[image.Point]Tile
	xMax, yMax int
}

func (pd PrintingDepartment) String() string {
	result := ""

	for y := range pd.yMax {
		for x := range pd.xMax {
			result += string(pd.grid[image.Pt(x, y)].Value)
		}

		result += "\n"
	}

	return result
}

func (pd PrintingDepartment) paperAccessibleByForklift(remove bool) int {
	result := 0

	markForRemoval := make(map[image.Point]struct{}, 0)

	for y := range pd.yMax {
		for x := range pd.xMax {
			currentPosition := image.Pt(x, y)

			// 1. Make sure tile exists in the grid
			// 2. Only count neighbours of paper, not of empty tiles
			if tileType, ok := pd.grid[currentPosition]; ok && tileType != Paper {
				continue
			}

			paperNeighbourCount := pd.countPaperNeighbours(currentPosition)

			slog.Debug("neighbours", slog.Int("x", x), slog.Int("y", y), slog.Int("paperNeighbourCount", paperNeighbourCount))

			if paperNeighbourCount < 4 {
				result++

				if remove {
					markForRemoval[currentPosition] = struct{}{}
				}
			}
		}
	}

	if remove {
		slog.Debug("marked for removal", slog.Int("count", len(markForRemoval)), slog.String("map", fmt.Sprintf("%+v", markForRemoval)))

		for key := range markForRemoval {
			pd.grid[key] = Empty
		}
	}

	return result
}

func (pd PrintingDepartment) countPaperNeighbours(position image.Point) int {
	result := 0

	for _, direction := range []image.Point{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}} {
		candidate := position.Add(direction)

		if tileType, ok := pd.grid[candidate]; ok && tileType == Paper {
			result++
		}
	}

	return result
}

func parseInput(input string) (PrintingDepartment, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := PrintingDepartment{
		grid: make(map[image.Point]Tile),
	}

	x, y := 0, 0

	for scanner.Scan() {
		xLine := strings.Split(scanner.Text(), " ")
		xLine = slices.DeleteFunc(xLine, func(s string) bool { return s == "" })

		if len(xLine) != 1 {
			return PrintingDepartment{}, fmt.Errorf("unexpected format on line '%+v'", xLine)
		}

		x = 0

		for _, r := range xLine[0] {
			newTile := Tiles.Parse(r)
			if newTile == nil {
				return PrintingDepartment{}, fmt.Errorf("unexpected rune '%s' on line '%s'", string(r), xLine[0])
			}

			result.grid[image.Pt(x, y)] = *newTile

			x++
		}

		y++
	}

	result.xMax = x
	result.yMax = y

	return result, nil
}
