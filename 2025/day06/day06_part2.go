// Package day06 for Advent of Code 2025, day 6, part 2.
// https://adventofcode.com/2025/day/6
package day06

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

func parseInputPart2(input string) (TrashCompactor, error) {
	xLines := strings.Split(input, "\n")

	if len(xLines) < 2 {
		return TrashCompactor{}, nil
	}

	operatorLine := xLines[len(xLines)-2]
	numberLines := xLines[:len(xLines)-2]

	slog.Debug("split lines", slog.String("slice", fmt.Sprintf("%#v", numberLines)), slog.String("operators", fmt.Sprintf("%+v", operatorLine)))

	result := TrashCompactor{
		problems: make([]problem, 0),
	}

	newProblem := problem{numbers: make([]int, 0)}

	for column := len(operatorLine) - 1; column >= 0; column-- {
		sb := strings.Builder{}

		for _, nl := range numberLines {
			slog.Debug("look down the column on each number line", slog.String("nl", nl), slog.Int("column", column), slog.String("rune", nl[column:column+1]))

			if nextChar := nl[column : column+1]; nextChar != " " {
				sb.WriteString(nextChar)
			}
		}

		sbs := strings.TrimSpace(sb.String())

		slog.Debug("rolled up number", slog.String("sbs", sbs))

		if len(sbs) == 0 {
			continue
		}

		npn, err := strconv.Atoi(sbs)
		if err != nil {
			return TrashCompactor{}, fmt.Errorf("converting number '%s': %w", sbs, err)
		}

		newProblem.numbers = append(newProblem.numbers, npn)

		slog.Debug("check operator after numbers", slog.String("ol", operatorLine), slog.Int("column", column), slog.String("rune", operatorLine[column:column+1]))

		if operator := operatorLine[column : column+1]; operator != " " {
			newProblem.operator = []rune(operator)[0]
			result.problems = append(result.problems, newProblem)

			newProblem = problem{numbers: make([]int, 0)}
		}
	}

	slog.Debug("done parsing", slog.Any("result", result))

	return result, nil
}

func Part2(input string) (int, error) {
	tc, err := parseInputPart2(input)
	if err != nil {
		return 0, err
	}

	result := tc.grandTotal()

	return result, nil
}
