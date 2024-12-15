// Package day15 for Advent of Code 2024, day 15, part 1.
// https://adventofcode.com/2024/day/15
package day15

import "image"

func Part1(input string) (int, error) {
	wh, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	for _, mi := range wh.moveInstructions {
		wh.execute(mi)
	}

	return wh.sumBoxGPS(), nil
}

func (wh Warehouse) findRobot() image.Point {
	for x := range wh.floorBounds.Max.X {
		for y := range wh.floorBounds.Max.Y {
			if wh.grid[image.Pt(x, y)] == Robot {
				return image.Pt(x, y)
			}
		}
	}

	panic("the robot has gone missing")
}

var moves = map[Move]image.Point{
	Up:    image.Pt(0, -1),
	Right: image.Pt(1, 0),
	Down:  image.Pt(0, 1),
	Left:  image.Pt(-1, 0),
}

func (wh Warehouse) findFirstEmpty(dir Move) (image.Point, bool) {
	for cursor := wh.findRobot(); cursor.In(wh.floorBounds); cursor = cursor.Add(moves[dir]) {
		switch wh.grid[cursor] {
		case Wall:
			return image.Point{}, false

		case Floor:
			return cursor, true

		case Robot, Box:
			continue
		}
	}

	return image.Point{}, false
}

func (wh Warehouse) execute(dir Move) {
	empty, canPush := wh.findFirstEmpty(dir)
	if !canPush {
		return
	}

	robot := wh.findRobot()

	for cursor := empty; cursor.In(wh.floorBounds) && !cursor.Eq(robot); cursor = cursor.Sub(moves[dir]) {
		wh.grid[cursor], wh.grid[cursor.Sub(moves[dir])] = wh.grid[cursor.Sub(moves[dir])], wh.grid[cursor]
	}
}

func (wh Warehouse) sumBoxGPS() int {
	result := 0

	for x := range wh.floorBounds.Max.X {
		for y := range wh.floorBounds.Max.Y {
			if wh.grid[image.Pt(x, y)] == Box {
				result += x + (y * 100)
			}
		}
	}

	return result
}
