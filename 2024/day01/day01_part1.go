package day01

import (
	"bufio"
	"bytes"
	"errors"
	"slices"
	"strconv"
	"strings"

	"go.jlucktay.dev/adventofcode/crunchy"
)

type twoLists struct {
	left  []int
	right []int
}

func parseInput(input string) (twoLists, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := twoLists{
		left:  make([]int, 0),
		right: make([]int, 0),
	}

	for scanner.Scan() {
		xLine := strings.Split(scanner.Text(), " ")
		xLine = slices.DeleteFunc(xLine, func(s string) bool { return s == "" })

		left, err := strconv.Atoi(xLine[0])
		if err != nil {
			return twoLists{}, err
		}

		right, err := strconv.Atoi(xLine[1])
		if err != nil {
			return twoLists{}, err
		}

		result.left = append(result.left, left)
		result.right = append(result.right, right)
	}

	return result, nil
}

func ListDistance(input string) (int, error) {
	lists, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	left := slices.Clone(lists.left)
	slices.Sort(left)
	right := slices.Clone(lists.right)
	slices.Sort(right)

	if len(left) != len(right) {
		return 0, errors.New("list lengths differ")
	}

	result := 0

	for index := range left {
		result += crunchy.AbsoluteDiff(left[index], right[index])
	}

	return result, nil
}
