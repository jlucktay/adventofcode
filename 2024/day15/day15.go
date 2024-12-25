// Package day15 for Advent of Code 2024, day 15.
// https://adventofcode.com/2024/day/15
package day15

import (
	"fmt"
	"image"
	"strings"

	"github.com/orsinium-labs/enum"
)

func parseInput(input string) (Warehouse, error) {
	if len(input) == 0 {
		return Warehouse{}, nil
	}

	result := Warehouse{
		floorBounds:      image.Rect(0, 0, 0, 0),
		grid:             make(map[image.Point]Tile),
		moveInstructions: make([]Move, 0),
	}

	lines := strings.Split(input, "\n")

	parsingTiles := true

	for y, line := range lines {
		if len(line) == 0 {
			parsingTiles = false
			continue
		}

		for x, letter := range line {
			if parsingTiles {
				parsed := Tiles.Parse(letter)
				if parsed == nil {
					return Warehouse{}, fmt.Errorf("parsing tile '%s'", string(letter))
				}

				result.grid[image.Pt(x, y)] = *parsed

				if x+1 > result.floorBounds.Max.X {
					result.floorBounds.Max.X = x + 1
				}
			} else {
				parsed := Moves.Parse(letter)
				if parsed == nil {
					return Warehouse{}, fmt.Errorf("parsing instruction '%s'", string(letter))
				}

				result.moveInstructions = append(result.moveInstructions, *parsed)
			}
		}

		if parsingTiles && y+1 > result.floorBounds.Max.Y {
			result.floorBounds.Max.Y = y + 1
		}
	}

	return result, nil
}

type Move enum.Member[rune]

var (
	Up    = Move{'^'}
	Right = Move{'>'}
	Down  = Move{'v'}
	Left  = Move{'<'}

	LeftBox  = Move(BoxLeft)
	RightBox = Move(BoxRight)

	Moves = enum.New(Up, Right, Down, Left, LeftBox, RightBox)
)

type Tile enum.Member[rune]

var (
	Robot = Tile{'@'}
	Floor = Tile{'.'}
	Wall  = Tile{'#'}
	Box   = Tile{'O'}

	BoxLeft  = Tile{'['}
	BoxRight = Tile{']'}

	Tiles = enum.New(Robot, Floor, Wall, Box)
)

type warehouseGrid map[image.Point]Tile

type Warehouse struct {
	floorBounds      image.Rectangle
	grid             warehouseGrid
	moveInstructions []Move
}

func (wh Warehouse) String() string {
	sb := strings.Builder{}

	sb.WriteString("Warehouse:\n")
	sb.WriteString(fmt.Sprintf("\tFloor bounds: %s\n", wh.floorBounds))
	sb.WriteString("\tMove instructions: ")

	for _, moveInst := range wh.moveInstructions {
		sb.WriteRune(moveInst.Value)
	}

	sb.WriteString("\n")

	for y := range wh.floorBounds.Max.Y {
		for x := range wh.floorBounds.Max.X {
			sb.WriteRune(wh.grid[image.Pt(x, y)].Value)
		}

		sb.WriteString("\n")
	}

	return sb.String()
}
