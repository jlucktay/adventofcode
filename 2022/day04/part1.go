package day04

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func SectionIDFullyContain(input string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	result := 0

	for scanner.Scan() {
		line := scanner.Text()

		xLine := strings.Split(line, ",")
		if len(xLine) != 2 {
			return 0, fmt.Errorf("line '%s' is not split by one comma only", line)
		}

		leftAssignment, err := parseAssignment(xLine[0])
		if err != nil {
			return 0, err
		}

		rightAssignment, err := parseAssignment(xLine[1])
		if err != nil {
			return 0, err
		}

		if assignmentsFullyContain(leftAssignment, rightAssignment) {
			result++
		}
	}

	return result, nil
}

func parseAssignment(input string) ([]int, error) {
	xInput := strings.Split(input, "-")
	if len(xInput) != 2 {
		return nil, fmt.Errorf("string '%s' is not split by one hyphen only", input)
	}

	left, err := strconv.ParseInt(xInput[0], 10, 32)
	if err != nil {
		return nil, err
	}

	right, err := strconv.ParseInt(xInput[1], 10, 32)
	if err != nil {
		return nil, err
	}

	return []int{int(left), int(right)}, nil
}

func assignmentsFullyContain(left, right []int) bool {
	if left[0] <= right[0] && left[1] >= right[1] {
		return true
	}

	if right[0] <= left[0] && right[1] >= left[1] {
		return true
	}

	return false
}
