// Package main for Advent of Code 2023, day 9, part 1
// https://adventofcode.com/2023/day/9
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// History is one puzzle input line.
type History []int

func (h History) String() string {
	sb := strings.Builder{}

	for i, v := range h {
		sb.WriteString(fmt.Sprintf("%d", v))

		if i < len(h)-1 {
			sb.WriteRune(' ')
		}
	}

	return sb.String()
}

// OASISReport is a collection of puzzle input lines.
type OASISReport []History

func (or OASISReport) String() string {
	sb := strings.Builder{}

	for _, h := range or {
		sb.WriteString(h.String())
		sb.WriteRune('\n')
	}

	return sb.String()
}

func parseInput(inputLines []string) (OASISReport, error) {
	result := make(OASISReport, 0)

	for ilIndex := range inputLines {
		if len(inputLines[ilIndex]) == 0 {
			continue
		}

		parsedLine, err := parseLine(inputLines[ilIndex])
		if err != nil {
			return nil, err
		}

		result = append(result, parsedLine)
	}

	return result, nil
}

func parseLine(inputLine string) (History, error) {
	if len(inputLine) == 0 {
		return nil, nil
	}

	result := make(History, 0)

	xLine := strings.Split(inputLine, " ")
	for _, num := range xLine {
		parsed, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}

		result = append(result, parsed)
	}

	return result, nil
}

func (h History) nextValue() int {
	lenH := len(h)
	if lenH == 0 {
		return 0
	}

	deltas := h.reduceToZeroCount()

	deltaSum := 0
	for i := 0; i < len(deltas); i++ {
		deltaSum += deltas[i]
	}

	return h[lenH-1] + deltaSum
}

// reduceToZeroCount finds how many rows it takes to get to all zeroes, and returns a list of differences between the
// second last and last elements on each line. The length of this list is the number of rows it took.
func (h History) reduceToZeroCount() []int {
	if len(h) == 0 {
		return nil
	}

	deltas := make([]int, 0)

	if lenH := len(h); lenH >= 2 {
		deltas = append(deltas, h[lenH-1]-h[lenH-2])
	}

	bottomLine := make(History, len(h))

	copy(bottomLine, h)

	for !bottomLine.allZeroes() {
		for i := 0; i < len(bottomLine)-1; i++ {
			thisValue := bottomLine[i]
			nextValue := bottomLine[i+1]

			bottomLine[i] = nextValue - thisValue
		}

		truncateOne := len(bottomLine) - 1
		bottomLine = bottomLine[:truncateOne]

		if lenBL := len(bottomLine); lenBL >= 2 {
			deltas = append(deltas, bottomLine[lenBL-1]-bottomLine[lenBL-2])
		}
	}

	truncateOneDelta := len(deltas) - 1
	deltas = deltas[:truncateOneDelta]

	return deltas
}

func (h History) allZeroes() bool {
	if len(h) == 0 {
		return false
	}

	if len(h) == 1 {
		return h[0] == 0
	}

	for i := 0; i < len(h); i++ {
		if h[i] != 0 {
			return false
		}
	}

	return true
}

func Part1(inputLines []string) (int, error) {
	puzzle, err := parseInput(inputLines)
	if err != nil {
		return 0, err
	}

	total := 0

	for i := 0; i < len(puzzle); i++ {
		next := puzzle[i].nextValue()

		total += next
	}

	return total, nil
}
