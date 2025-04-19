// Package day14 for Advent of Code 2024, day 14.
// https://adventofcode.com/2024/day/14
package day14

import (
	"fmt"
	"image"
	"strings"
)

func parseInput(input string) (Robots, error) {
	lines := strings.Split(input, "\n")

	result := Robots{
		bots:   make([]*Robot, 0),
		bounds: image.Rectangle{},
	}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		var pX, pY, vX, vY int

		n, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &pX, &pY, &vX, &vY)
		if err != nil || n != 4 {
			return Robots{}, fmt.Errorf("scanning: %w", err)
		}

		result.bots = append(result.bots, &Robot{
			position: image.Point{pX, pY},
			velocity: image.Point{vX, vY},
		})
	}

	return result, nil
}
