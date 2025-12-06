// Package day06 for Advent of Code 2025, day 6.
// https://adventofcode.com/2025/day/6
package day06

import (
	"bufio"
	"bytes"
	"fmt"
	"log/slog"
	"slices"
	"strings"
)

type TrashCompactor struct {
	problems []problem
}

type problem struct {
	numbers  []int
	operator rune
}

func (p problem) solve() int {
	result := 0

	switch p.operator {
	case '+':
		for _, number := range p.numbers {
			result += number
		}

	case '*':
		result = p.numbers[0]

		for i := 1; i < len(p.numbers); i++ {
			result *= p.numbers[i]
		}

	default:
		slog.Error("unknown operator", slog.String("rune", string(p.operator)))
	}

	return result
}

func (tc TrashCompactor) grandTotal() int {
	result := 0

	for _, p := range tc.problems {
		result += p.solve()
	}

	return result
}

func parseInput(input string) (TrashCompactor, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := TrashCompactor{}

	tmpLines := [][]int{}
	tmpLine := 0
	var operatorsLine []string

	for scanner.Scan() {
		xLine := strings.Split(scanner.Text(), " ")
		xLine = slices.DeleteFunc(xLine, func(s string) bool { return s == "" })

		tmpLines = append(tmpLines, make([]int, 0))

		for _, rawNumber := range xLine {
			if len(xLine) >= 1 && (xLine[0] == "*" || xLine[0] == "+") {
				operatorsLine = xLine
				break
			}

			var number int

			numScanned, err := fmt.Sscanf(rawNumber, "%d", &number)
			if numScanned != 1 {
				return TrashCompactor{}, fmt.Errorf("incorrect number of fields scanned: %d", numScanned)
			}
			if err != nil {
				return TrashCompactor{}, fmt.Errorf("scanning '%s': %w", rawNumber, err)
			}

			tmpLines[tmpLine] = append(tmpLines[tmpLine], number)
		}

		tmpLine++
	}

	slog.Debug("done scanning", slog.String("tmpLines", fmt.Sprintf("%+v", tmpLines)))

	if len(tmpLines) >= 2 {
		tmpLines = tmpLines[:len(tmpLines)-1]
		slog.Debug("done trimming", slog.String("tmpLines", fmt.Sprintf("%+v", tmpLines)))
	} else {
		return TrashCompactor{}, nil
	}

	slog.Debug("line with operators", slog.String("operatorsLine", fmt.Sprintf("%+v", operatorsLine)))

	lineLen := len(tmpLines[0])

	for i := 1; i < len(tmpLines); i++ {
		if len(tmpLines[i]) != lineLen {
			return TrashCompactor{}, fmt.Errorf("incorrect length on line %d: %+v", i+1, tmpLines[i])
		}
	}

	if len(operatorsLine) != lineLen {
		return TrashCompactor{}, fmt.Errorf("incorrect operators line length: %+v", operatorsLine)
	}

	result.problems = make([]problem, 0)

	for column := range lineLen {
		newProblem := problem{numbers: make([]int, 0)}

		for _, numbers := range tmpLines {
			newProblem.numbers = append(newProblem.numbers, numbers[column])
		}

		newProblem.operator = []rune(operatorsLine[column])[0]

		result.problems = append(result.problems, newProblem)
	}

	slog.Debug("problems", slog.String("slice", fmt.Sprintf("%+v", result.problems)))

	return result, nil
}
