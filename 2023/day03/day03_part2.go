// Package main for Advent of Code 2023, day 3, part 2
// https://adventofcode.com/2023/day/3
package main

import (
	"log/slog"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func Part2(inputLines []string) (int, error) {
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

		for rowIndex < len(row) {
			switch row[rowIndex] {
			case '*':
				slog.Debug("found a gear",
					slog.Int("row", schematicRow),
					slog.Int("column", rowIndex))

				gearComponents, err := lookAroundGear(schematic, schematicRow, rowIndex)
				if err != nil {
					return 0, err
				}

				result += gearComponents[0] * gearComponents[1]

			default:
			}

			rowIndex++
		}

		schematicRow++
	}

	return result, nil
}

func lookAroundGear(schematic [][]rune, row, column int) ([2]int, error) {
	// Make sure we have at least one row before proceeding.
	if row < 0 || row >= len(schematic) {
		return [2]int{}, nil
	}

	result := [2]int{}

	// Check the row above, if it exists.
	if row-1 >= 0 {
		for i := column - 1; i <= column+1; i++ {
			if i >= 0 && i < len(schematic[row-1]) {
				if unicode.IsDigit(schematic[row-1][i]) {
					slog.Debug("number nearby",
						slog.Int("row", row-1),
						slog.Int("column", i),
						slog.String("digit", string(schematic[row-1][i])))

					entireNumber, err := discoverEntireNumber(schematic, row-1, i)
					if err != nil {
						return [2]int{}, err
					}

					if result[0] == 0 {
						result[0] = entireNumber
					} else if result[1] == 0 && result[0] != entireNumber {
						result[1] = entireNumber
						return result, nil
					}
				}
			}
		}
	}

	// Check the same row, before and after, if we're not at the start/end of the row.
	if column-1 >= 0 {
		if unicode.IsDigit(schematic[row][column-1]) {
			slog.Debug("number nearby",
				slog.Int("row", row),
				slog.Int("column", column-1),
				slog.String("digit", string(schematic[row][column-1])))

			entireNumber, err := discoverEntireNumber(schematic, row, column-1)
			if err != nil {
				return [2]int{}, err
			}

			if result[0] == 0 {
				result[0] = entireNumber
			} else if result[1] == 0 && result[0] != entireNumber {
				result[1] = entireNumber
				return result, nil
			}
		}
	}

	if column+1 < len(schematic[row]) {
		if unicode.IsDigit(schematic[row][column+1]) {
			slog.Debug("number nearby",
				slog.Int("row", row),
				slog.Int("column", column+1),
				slog.String("digit", string(schematic[row][column+1])))

			entireNumber, err := discoverEntireNumber(schematic, row, column+1)
			if err != nil {
				return [2]int{}, err
			}

			if result[0] == 0 {
				result[0] = entireNumber
			} else if result[1] == 0 && result[0] != entireNumber {
				result[1] = entireNumber
				return result, nil
			}
		}
	}

	// Check the row below, if it exists.
	if row+1 < len(schematic) {
		for i := column - 1; i <= column+1; i++ {
			if i >= 0 && i < len(schematic[row+1]) {
				if unicode.IsDigit(schematic[row+1][i]) {
					slog.Debug("number nearby",
						slog.Int("row", row+1),
						slog.Int("column", i),
						slog.String("digit", string(schematic[row+1][i])))

					entireNumber, err := discoverEntireNumber(schematic, row+1, i)
					if err != nil {
						return [2]int{}, err
					}

					if result[0] == 0 {
						result[0] = entireNumber
					} else if result[1] == 0 && result[0] != entireNumber {
						result[1] = entireNumber
						return result, nil
					}
				}
			}
		}
	}

	return [2]int{}, nil
}

func discoverEntireNumber(schematic [][]rune, row, column int) (int, error) {
	emptyString := ""

	if len(schematic) <= row {
		return 0, nil
	}

	// walk to the left
	if column >= 0 && column < len(schematic[row]) {
		for i := column; i >= 0; i-- {
			if unicode.IsDigit(schematic[row][i]) {
				emptyString = string(schematic[row][i]) + emptyString
			} else {
				break
			}
		}
	}

	// walk to the right
	if column+1 < len(schematic[row]) {
		for i := column + 1; i < len(schematic[row]); i++ {
			if unicode.IsDigit(schematic[row][i]) {
				emptyString = emptyString + string(schematic[row][i])
			} else {
				break
			}
		}
	}

	number, err := strconv.Atoi(emptyString)
	if err != nil {
		return 0, err
	}

	return number, nil
}
