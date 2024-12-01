package day01

import (
	"bufio"
	"bytes"
	"errors"
	"slices"
	"strconv"
	"strings"
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

	if len(result.left) != len(result.right) {
		return twoLists{}, errors.New("list lengths differ")
	}

	return result, nil
}
