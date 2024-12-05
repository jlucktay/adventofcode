// Package day05 for Advent of Code 2024, day 5, part 2.
// https://adventofcode.com/2024/day/5
package day05

import "slices"

func Part2(input string) (int, error) {
	mp, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := 0

	for _, upd := range mp.updates {
		if !mp.updateInRightOrder(upd) {
			slices.SortFunc(upd, func(a, b int) int {
				if mp.checkPageOrder(a, b) {
					return 1
				} else {
					return -1
				}
			})

			result += upd.middlePageNumber()
		}
	}

	return result, nil
}
