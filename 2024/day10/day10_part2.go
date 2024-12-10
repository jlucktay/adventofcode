// Package day10 for Advent of Code 2024, day 10, part 2.
// https://adventofcode.com/2024/day/10
package day10

import (
	"slices"
	"strings"
)

func Part2(input string) (int, error) {
	hm, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := 0

	for _, th := range hm.trailheads {
		for _, p := range hm.peaks {
			candidates := make([]string, 0)

			shortests, err := hm.graph.ShortestAll(th, p)
			if err != nil {
				continue
			}

			for _, shortest := range shortests.Paths {
				candidates = append(candidates, strings.Join(shortest, "->"))
			}

			longests, err := hm.graph.LongestAll(th, p)
			if err != nil {
				continue
			}

			for _, longest := range longests.Paths {
				candidates = append(candidates, strings.Join(longest, "->"))
			}

			slices.Sort(candidates)
			result += len(slices.Compact(candidates))
		}
	}

	return result, nil
}
