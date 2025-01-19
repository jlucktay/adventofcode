// Package day18 for Advent of Code 2024, day 18, part 2.
// https://adventofcode.com/2024/day/18
package day18

import (
	"errors"
	"fmt"
	"image"
	"strings"

	"github.com/RyanCarrier/dijkstra/v2"
)

func Part2(input string) (string, error) {
	rr, err := parseInput(input)
	if err != nil {
		return "", err
	}

	if len(rr.grid) == 0 {
		return "", nil
	}

	for _, urb := range rr.undroppedRawBytes {
		var x, y int

		n, err := fmt.Sscanf(urb, "%d,%d", &x, &y)
		if err != nil || n != 2 {
			return "", fmt.Errorf("scanning undropped raw byte '%s' parsed %d items: %w", urb, n, err)
		}

		newTile := image.Pt(x, y)

		// Set new tile in grid, and remove vertex from graph.
		rr.grid[newTile] = Byte

		if err := rr.graph.RemoveVertexAndArcs(newTile); err != nil {
			return "", fmt.Errorf("removing vertex '%s' and arcs: %w", newTile, err)
		}

		// Retry pathfinding, and if we get a 'no path found' error, return the new tile in a correctly-formatted string.
		if _, err = rr.MinimumStepsToExit(); err != nil {
			if errors.Is(err, dijkstra.ErrNoPath) {
				return strings.Trim(newTile.String(), "()"), nil
			}

			return "", fmt.Errorf("minimum steps to exit: %w", err)
		}
	}

	return "", errors.New("no blocking tile found")
}
