package day04

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func sectionIDAnalysis(analyser func(left, right []int) bool, input string) (int, error) {
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

		if analyser(leftAssignment, rightAssignment) {
			result++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scanning input: %v", err)
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
		return nil, fmt.Errorf("parsing integer: %w", err)
	}

	right, err := strconv.ParseInt(xInput[1], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("parsing integer: %w", err)
	}

	return []int{int(left), int(right)}, nil
}
