// Package day12 for Advent of Code 2024, day 12, part 2.
// https://adventofcode.com/2024/day/12
package day12

import "image"

func Part2(input string) (int, error) {
	g, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	return g.sides(), nil
}

func (g Garden) sides() int {
	seen := map[image.Point]bool{}
	totalSides := 0

	for point := range g {
		if seen[point] {
			continue
		}

		seen[point] = true

		area := 1
		sides := 0
		queue := []image.Point{point}

		for len(queue) > 0 {
			queuePoint := queue[0]
			queue = queue[1:]

			for _, direction := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				if neighbour := queuePoint.Add(direction); g[neighbour] != g[queuePoint] {
					rotated := queuePoint.Add(image.Point{-direction.Y, direction.X})
					if g[rotated] != g[queuePoint] || g[rotated.Add(direction)] == g[queuePoint] {
						sides++
					}
				} else if !seen[neighbour] {
					seen[neighbour] = true
					queue = append(queue, neighbour)
					area++
				}
			}
		}

		totalSides += area * sides
	}

	return totalSides
}
