// Package day15 for Advent of Code 2024, day 15, part 2.
// https://adventofcode.com/2024/day/15
package day15

import (
	"image"
	"slices"
	"strconv"
)

func Part2(input string) (int, error) {
	wh, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	wwh := wh.widen()

	for _, mi := range wwh.moveInstructions {
		wwh.execute(mi)
	}

	return wwh.sumBoxGPS(), nil
}

type WideWarehouse struct {
	Warehouse
}

func (wh Warehouse) widen() WideWarehouse {
	result := WideWarehouse{
		Warehouse: Warehouse{
			floorBounds: image.Rectangle{
				Min: image.Point{X: 0, Y: 0},
				Max: image.Point{
					X: wh.floorBounds.Max.X * 2,
					Y: wh.floorBounds.Max.Y,
				},
			},
			grid:             make(warehouseGrid),
			moveInstructions: slices.Clip(slices.Clone(wh.moveInstructions)),
		},
	}

	// Range over the original un-widened warehouse and extrapolate out into the widened warehouse's grid.
	for y := range wh.floorBounds.Max.Y {
		for x := range wh.floorBounds.Max.X {
			newLeft := image.Pt(2*x, y)
			newRight := image.Pt((2*x)+1, y)

			switch wh.grid[image.Pt(x, y)] {
			case Wall:
				result.grid[newLeft] = Wall
				result.grid[newRight] = Wall

			case Box:
				result.grid[newLeft] = BoxLeft
				result.grid[newRight] = BoxRight

			case Floor:
				result.grid[newLeft] = Floor
				result.grid[newRight] = Floor

			case Robot:
				result.grid[newLeft] = Robot
				result.grid[newRight] = Floor
			}
		}
	}

	// Sanity check on the widened warehouse.
	tileCounts := make(map[Tile]int)

	for _, tile := range result.grid {
		tileCounts[tile]++
	}

	if tileCounts[Robot] != 1 && len(result.grid) > 0 {
		panic("should have one robot only and not " + strconv.Itoa(tileCounts[Robot]))
	}

	if tileCounts[Box] != 0 {
		panic("should not have any single-tile boxes")
	}

	return result
}

func (wwh WideWarehouse) execute(dir Move) {
	// Start searching from the robot's current position.
	robot := wwh.findRobot()
	queue := []image.Point{robot}
	boxes := make(map[image.Point]Tile)

	for len(queue) > 0 {
		// Progress the queue.
		cursor := queue[0]
		queue = queue[1:]

		// Skip going around the loop if we've already seen this box here at these coordinates.
		if _, box := boxes[cursor]; box {
			continue
		}

		// Copy the tile type at the current cursor into the working map of boxes to (maybe) move around later.
		boxes[cursor] = wwh.grid[cursor]

		switch n := cursor.Add(moves[dir]); wwh.grid[n] {
		case Wall:
			// If we hit a wall, we're done, so don't make any changes.
			return

		case BoxLeft, BoxRight:
			// Add this half of the 2-wide box.
			queue = append(queue, n)

			// Add the other half of the 2-wide box.
			m := Moves.Parse(wwh.grid[n].Value)
			queue = append(queue, n.Add(moves[*m]))
		}
	}

	// Temporarily set all tiles where the boxes used to sit to empty floor tiles.
	for b := range boxes {
		wwh.grid[b] = Floor
	}

	// Add the move delta to the tiles where the boxes used to be, and set those tiles to the moved box half type.
	for b := range boxes {
		wwh.grid[b.Add(moves[dir])] = boxes[b]
	}

	// Update the robot's location.
	wwh.grid[robot.Add(moves[dir])] = Robot
	wwh.grid[robot] = Floor
}
