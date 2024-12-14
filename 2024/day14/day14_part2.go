// Package day14 for Advent of Code 2024, day 14, part 2.
// https://adventofcode.com/2024/day/14
package day14

import (
	"fmt"
	"image"
)

func Part2(input string, bounds image.Rectangle, seconds int) (int, error) {
	r, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	r.bounds = bounds

	for sec := range seconds {
		for _, robot := range r.bots {
			robot.advance(1, r.bounds)
		}

		if r.botsAllDistinct() {
			fmt.Printf("%s\n", r)

			return sec + 1, nil
		}
	}

	return 0, nil
}

func (r Robots) botsAllDistinct() bool {
	foundBots := make(map[image.Point]int)

	for _, bot := range r.bots {
		foundBots[image.Point{bot.position.X, bot.position.Y}]++
	}

	for _, count := range foundBots {
		if count != 1 {
			return false
		}
	}

	return true
}
