package day01

import (
	"slices"

	"go.jlucktay.dev/adventofcode/crunchy"
)

func ListDistance(input string) (int, error) {
	lists, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	left := slices.Clone(lists.left)
	slices.Sort(left)
	right := slices.Clone(lists.right)
	slices.Sort(right)

	result := 0

	for index := range left {
		result += crunchy.AbsDiff(left[index], right[index])
	}

	return result, nil
}
