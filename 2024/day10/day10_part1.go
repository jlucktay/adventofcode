// Package day10 for Advent of Code 2024, day 10, part 1.
// https://adventofcode.com/2024/day/10
package day10

func Part1(input string) (int, error) {
	hm, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := 0

	for _, th := range hm.trailheads {
		for _, p := range hm.peaks {
			if _, err := hm.graph.Shortest(th, p); err != nil {
				continue
			}

			result++
		}
	}

	return result, nil
}
