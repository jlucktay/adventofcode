// Package day05 for Advent of Code 2024, day 5, part 1.
// https://adventofcode.com/2024/day/5
package day05

func Part1(input string) (int, error) {
	mp, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := 0

	for _, upd := range mp.updates {
		if mp.updateInRightOrder(upd) {
			result += upd.middlePageNumber()
		}
	}

	return result, nil
}

type rule [2]int

type update []int

type manualPrinter struct {
	rules   []rule
	updates []update
}

func (mp manualPrinter) updateInRightOrder(upd update) bool {
	for outerIndex, firstPageNumber := range upd {
		for _, secondPageNumber := range upd[outerIndex+1:] {
			if !mp.checkPageOrder(firstPageNumber, secondPageNumber) {
				return false
			}
		}
	}

	return true
}

func (mp manualPrinter) checkPageOrder(before, after int) bool {
	for _, rul := range mp.rules {
		if rul[0] == before && rul[1] == after {
			return true
		}
	}

	return false
}

func (u update) middlePageNumber() int {
	// Even number of elements == barf
	if len(u)%2 == 0 {
		return 0
	}

	return u[len(u)/2]
}
