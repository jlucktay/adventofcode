// Package day12 for Advent of Code 2024, day 12, part 1.
// https://adventofcode.com/2024/day/12
package day12

import "image"

func Part1(input string) (int, error) {
	g, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	return g.perimeters(), nil
}

func (g Garden) perimeters() int {
	seen := make(map[image.Point]bool)
	totalPerimiter := 0

	for point := range g {
		if seen[point] {
			continue
		}

		seen[point] = true

		area := 1
		perimeter := 0
		queue := []image.Point{point}

		for len(queue) > 0 {
			queuePoint := queue[0]
			queue = queue[1:]

			for _, direction := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				if neighbour := queuePoint.Add(direction); g[neighbour] != g[queuePoint] {
					perimeter++
				} else if !seen[neighbour] {
					seen[neighbour] = true
					queue = append(queue, neighbour)
					area++
				}
			}
		}

		totalPerimiter += area * perimeter
	}

	return totalPerimiter
}
