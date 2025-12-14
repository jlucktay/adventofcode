// Package day08 for Advent of Code 2025, day 8, part 2.
// https://adventofcode.com/2025/day/8
package day08

import (
	"log/slog"
	"slices"

	"go.jlucktay.dev/adventofcode/crunchy/slice"
)

// Part2 was demystified for me by the following, and a big 'thank you' goes out to the authors of both:
//   - [This summary] of what to tackle the problem with
//   - [This write-up] of what Kruskal's algorithm is used for: constructing a minimum spanning tree of a weighted connected undirected graph
//
// [This summary]: https://github.com/erik-adelbert/aoc/blob/main/2025/README.md#day-8-playground-
// [This write-up]: https://gist.github.com/DanilAndreev/e519d77eff91f03f09616c9170db7941
func Part2(input string) (int, error) {
	pg, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	if len(pg.boxes) == 0 {
		return 0, nil
	}

	slog.Debug("calculating distances between boxes", slog.Int("boxes", len(pg.boxes)))
	pg.calculateDistances()
	slog.Debug("calculated distances between boxes", slog.Int("boxes", len(pg.boxes)), slog.Int("distancePairs", len(pg.distances)))

	// At this point, we have the list of edges sorted by weight/distance in 'pg.distances'.

	// Seed the first circuit with the two boxes with the shortest distance between.
	circuits := make([][]*junctionBox, 0)
	circuits = append(circuits, []*junctionBox{pg.distances[0].left, pg.distances[0].right})

	// Pop the shortest distance pair off the front of that slice.
	pg.distances = pg.distances[1:]

	var lastTwoBoxes int
	circuitSizeCutoff := len(pg.boxes)

	// Range across all subsequent edges (in ascending weight/distance order) after the already-removed lightest/shortest first element.
	for _, edge := range pg.distances {
		startConnection, endConnection := -1, -1

		// Look for start point occurrences in the circuits, and get the index of the circuit.
		for i, circuit := range circuits {
			if slices.ContainsFunc(circuit, func(jb *junctionBox) bool { return jb == edge.left }) {
				startConnection = i

				break
			}
		}

		// Look for end point occurrences in the circuits, and get the index of the circuit.
		for j, circuit := range circuits {
			if slices.ContainsFunc(circuit, func(jb *junctionBox) bool { return jb == edge.right }) {
				endConnection = j

				break
			}
		}

		// If the edge makes a cycle, skip it.
		if startConnection >= 0 && startConnection == endConnection {
			continue
		}

		// If edge's first point was found, add the edge's second point to the same circuit.
		if startConnection >= 0 {
			circuits[startConnection] = append(circuits[startConnection], edge.right)
		}

		// If edge's second point was found, add the edge's first point to the same circuit.
		if endConnection >= 0 {
			circuits[endConnection] = append(circuits[endConnection], edge.left)
		}

		// If edge is not already connected with any circuits, add a new circuit.
		if startConnection == -1 && endConnection == -1 {
			circuits = append(circuits, []*junctionBox{edge.left, edge.right})
		}

		// If edge connects two circuits, merge them.
		if startConnection >= 0 && endConnection >= 0 {
			newCircuit := slice.Union(circuits[startConnection], circuits[endConnection])
			circuits = append(circuits, newCircuit)

			circuits = slice.Delete(circuits, startConnection, endConnection)
		}

		// Check the size of all circuits connected so far; once we have one containing every junction box, that's our cutoff, so we don't need to keep working.
		var biggestCircuit int

		for _, circuit := range circuits {
			if lenCircuit := len(circuit); lenCircuit >= biggestCircuit {
				biggestCircuit = lenCircuit
			}
		}

		if biggestCircuit >= circuitSizeCutoff {
			slog.Debug("all boxes are now together in one circuit; last two connected boxes", slog.Any("left", edge.left), slog.Any("right", edge.right), slog.Int("distance", edge.distance))

			lastTwoBoxes = edge.left.coords[jbX] * edge.right.coords[jbX]

			break
		}
	}

	return lastTwoBoxes, nil
}
