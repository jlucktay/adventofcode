// Package day14 for Advent of Code 2024, day 14, part 1.
// https://adventofcode.com/2024/day/14
package day14

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/orsinium-labs/enum"
)

func Part1(input string, bounds image.Rectangle, seconds int) (int, error) {
	r, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	r.bounds = bounds

	for _, robot := range r.bots {
		robot.advance(seconds, r.bounds)
	}

	return r.safetyFactor(bounds.Dx() < 100), nil
}

type Robot struct {
	position, velocity image.Point
}

type Robots struct {
	bots   []*Robot
	bounds image.Rectangle
}

func (r *Robot) advance(seconds int, bounds image.Rectangle) {
	r.position.X += (r.velocity.X * seconds)
	r.position.Y += (r.velocity.Y * seconds)

	r.position = r.position.Mod(bounds)
}

func (r Robots) String() string {
	sb := strings.Builder{}

	sb.WriteString("Bot positions:\n")

	for index, bot := range r.bots {
		sb.WriteString(fmt.Sprintf("\t#%04d @ %d,%d with velocity %d,%d\n",
			index, bot.position.X, bot.position.Y, bot.velocity.X, bot.velocity.Y))
	}

	sb.WriteString("Bounds: " + r.bounds.String() + "\n")

	for y := range r.bounds.Max.Y {
		for x := range r.bounds.Max.X {
			robotsHere := 0

			for _, robot := range r.bots {
				if robot.position.X == x && robot.position.Y == y {
					robotsHere++
				}
			}

			if robotsHere == 0 {
				sb.WriteRune('.')
			} else if robotsHere >= 10 {
				sb.WriteRune('X')
			} else {
				sb.WriteString(strconv.Itoa(robotsHere))
			}
		}

		sb.WriteRune('\n')
	}

	return sb.String()
}

type Quadrant enum.Member[image.Rectangle]

var (
	// 11 wide
	// 7 tall

	TestTopLeft = Quadrant{image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{5, 3},
	}}
	TestTopRight = Quadrant{image.Rectangle{
		Min: image.Point{6, 0},
		Max: image.Point{11, 3},
	}}
	TestBottomLeft = Quadrant{image.Rectangle{
		Min: image.Point{0, 4},
		Max: image.Point{5, 7},
	}}
	TestBottomRight = Quadrant{image.Rectangle{
		Min: image.Point{6, 4},
		Max: image.Point{11, 7},
	}}

	TestQuadrants = enum.New(TestTopLeft, TestTopRight, TestBottomLeft, TestBottomRight)

	// 101 wide
	// 103 tall

	LiveTopLeft = Quadrant{image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{50, 51},
	}}
	LiveTopRight = Quadrant{image.Rectangle{
		Min: image.Point{51, 0},
		Max: image.Point{101, 51},
	}}
	LiveBottomLeft = Quadrant{image.Rectangle{
		Min: image.Point{0, 52},
		Max: image.Point{50, 103},
	}}
	LiveBottomRight = Quadrant{image.Rectangle{
		Min: image.Point{51, 52},
		Max: image.Point{101, 103},
	}}

	LiveQuadrants = enum.New(LiveTopLeft, LiveTopRight, LiveBottomLeft, LiveBottomRight)
)

func (r Robots) safetyFactor(test bool) int {
	quadrantCounts := make(map[Quadrant]int)

	var quadrants []Quadrant
	if test {
		quadrants = TestQuadrants.Members()
	} else {
		quadrants = LiveQuadrants.Members()
	}

	for _, robot := range r.bots {
		for _, q := range quadrants {
			if robot.position.In(q.Value) {
				quadrantCounts[q]++

				break
			}
		}
	}

	if len(quadrantCounts) == 0 {
		return 0
	}

	result := 1

	for _, count := range quadrantCounts {
		result *= count
	}

	return result
}
