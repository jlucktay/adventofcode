// Package main for Advent of Code 2023, day 3, part 1
// https://adventofcode.com/2023/day/3
package main

import (
	"log/slog"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Part1(inputLines []string) (int, error) {
	schematic := [][]rune{}
	schematicRow := 0

	for _, inputLine := range inputLines {
		schematic = append(schematic, make([]rune, 0))

		lineIndex := 0
		inputLineBytes := []byte(inputLine)

		for lineIndex < len(inputLineBytes) {
			r, size := utf8.DecodeRune(inputLineBytes[lineIndex:])

			schematic[schematicRow] = append(schematic[schematicRow], r)

			lineIndex += size
		}

		schematicRow++
	}

	schematicRow = 0
	result := 0

	for schematicRow < len(schematic) {
		row := schematic[schematicRow]
		rowIndex := 0
		onNumbers := false
		accumNumber := strings.Builder{}

		for rowIndex < len(row) {
			switch row[rowIndex] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				onNumbers = true
				accumNumber.WriteRune(row[rowIndex])

			default:
				if onNumbers {
					onNumbers = false
					accumNumberStr := accumNumber.String()
					accumNumber = strings.Builder{}

					accNum, err := strconv.Atoi(accumNumberStr)
					if err != nil {
						return 0, err
					}

					if lookAround(schematic, schematicRow, rowIndex-len(accumNumberStr), len(accumNumberStr)) {
						result += accNum

						slog.Debug("Part1 accumulating",
							slog.Int("accNum", accNum),
							slog.Int("result", result))
					} else {
						slog.Debug("Part1 not accumulating",
							slog.Int("accNum", accNum),
							slog.Int("result", result))
					}
				}
			}

			rowIndex++
		}

		if accumNumber.Len() > 0 {
			accumNumberStr := accumNumber.String()

			accNum, err := strconv.Atoi(accumNumberStr)
			if err != nil {
				return 0, err
			}

			if lookAround(schematic, schematicRow, rowIndex-len(accumNumberStr), len(accumNumberStr)) {
				result += accNum

				slog.Debug("Part1 accumulating",
					slog.Int("accNum", accNum),
					slog.Int("result", result))
			} else {
				slog.Debug("Part1 not accumulating",
					slog.Int("accNum", accNum),
					slog.Int("result", result))
			}
		}

		schematicRow++
	}

	return result, nil
}

func lookAround(schematic [][]rune, row, index, length int) bool {
	// Make sure we have at least one row before proceeding.
	if row < 0 || row >= len(schematic) {
		return false
	}

	slog.Debug("lookAround",
		slog.Int("len_schematic", len(schematic)),
		slog.Any("schematic", schematic),
		slog.Int("row", row),
		slog.Any("schematic_row", schematic[row]))

	// Check the row above, if it exists.
	if row-1 >= 0 {
		slog.Debug("checking row above",
			slog.String("row", string(schematic[row-1])),
			slog.Int("index", index),
			slog.Any("schematic[row-1]", schematic[row-1]))

		for i := index - 1; i <= index+length; i++ {
			slog.Debug("iterating over row above",
				slog.Int("index", i))

			if i >= 0 && i < len(schematic[row-1]) {
				slog.Debug("iterating over row above",
					slog.String("rune", string(schematic[row-1][i])))

				if isNotDotOrSymbol(schematic[row-1][i]) {
					return true
				}
			}
		}
	}

	// Check the same row, before and after, if we're not at the start/end of the row.
	if index-1 >= 0 {
		if isNotDotOrSymbol(schematic[row][index-1]) {
			return true
		}
	}

	if index+length < len(schematic[row]) {
		if isNotDotOrSymbol(schematic[row][index+length]) {
			return true
		}
	}

	// Check the row below, if it exists.
	if row+1 < len(schematic) {
		slog.Debug("checking row below",
			slog.String("row", string(schematic[row+1])),
			slog.Int("index", index),
			slog.Any("schematic[row+1]", schematic[row+1]))

		for i := index - 1; i <= index+length; i++ {
			slog.Debug("iterating over row below",
				slog.Int("index", i))

			if i >= 0 && i < len(schematic[row+1]) {
				slog.Debug("iterating over row below",
					slog.String("rune", string(schematic[row+1][i])))

				if isNotDotOrSymbol(schematic[row+1][i]) {
					return true
				}
			}
		}
	}

	return false
}

func isNotDotOrSymbol(r rune) bool {
	if r != '.' &&
		r != '0' &&
		r != '1' &&
		r != '2' &&
		r != '3' &&
		r != '4' &&
		r != '5' &&
		r != '6' &&
		r != '7' &&
		r != '8' &&
		r != '9' {
		return true
	}

	return false
}
