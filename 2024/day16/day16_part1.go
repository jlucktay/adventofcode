// Package day16 for Advent of Code 2024, day 16, part 1.
// https://adventofcode.com/2024/day/16
package day16

import (
	"cmp"
	"image"
	"maps"
	"math"
	"slices"
)

func Part1(input string) (int, error) {
	rm, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	if len(rm.grid) == 0 {
		return 0, nil
	}

	return rm.lowestScore(false), nil
}

type Status struct {
	position, direction image.Point
}

type QueueEntity struct {
	status Status
	cost   int
	path   map[image.Point]struct{}
}

func (rm ReindeerMaze) lowestScore(part2 bool) int {
	// Track costs as we go.
	distances := make(map[Status]int)

	// Set up a queue to process.
	queue := []QueueEntity{{
		Status{rm.start, image.Point{1, 0}},
		0,
		map[image.Point]struct{}{rm.start: {}},
	}}

	// Keep the lowest result.
	result := math.MaxInt

	// Count tiles on the best paths through the maze.
	bestSeat := make(map[image.Point]struct{})

	for len(queue) > 0 {
		// Move the lowest costs to the front of the queue.
		slices.SortFunc(queue, func(a, b QueueEntity) int { return cmp.Compare(a.cost, b.cost) })

		// Progress the queue.
		cursor := queue[0]
		queue = queue[1:]

		// Skip this iteration if there is already a lower cost recorded.
		if cost, ok := distances[cursor.status]; ok && cost < cursor.cost {
			continue
		}

		// Record the cost from this position/direction combination.
		distances[cursor.status] = cursor.cost

		if rm.grid[cursor.status.position] == End && result >= cursor.cost {
			result = cursor.cost

			maps.Copy(bestSeat, cursor.path)
		}

		// Iterate through candidates for the next move.
		for dir, cost := range map[image.Point]int{
			cursor.status.direction:                                 1,    // Continue straight.
			{-cursor.status.direction.Y, cursor.status.direction.X}: 1001, // Turn left.
			{cursor.status.direction.Y, -cursor.status.direction.X}: 1001, // Turn right.
		} {
			next := Status{cursor.status.position.Add(dir), dir}

			if rm.grid[next.position] == Wall {
				continue
			}

			// Make the next queue entry, incorporating the additional cost and the extra step in the explored path.
			path := maps.Clone(cursor.path)
			path[next.position] = struct{}{}

			queue = append(queue, QueueEntity{
				next,
				cursor.cost + cost,
				path,
			})
		}
	}

	if !part2 {
		return result
	}

	return len(bestSeat)
}
