// Package day02 for Advent of Code 2016, day 2.
// https://adventofcode.com/2016/day/2
package day02

import (
	"fmt"
	"strings"

	"github.com/orsinium-labs/enum"
)

type Direction enum.Member[rune]

var (
	Up    = Direction{'U'}
	Right = Direction{'R'}
	Down  = Direction{'D'}
	Left  = Direction{'L'}

	Directions = enum.New(Up, Right, Down, Left)
)

func (d Direction) String() string { return string(d.Value) }

type DigitInstruction []Direction

type Instructions []DigitInstruction

func parseInput(input string) (Instructions, error) {
	result := make(Instructions, 0)

	lines := strings.Fields(input)

	for lineNumber, line := range lines {
		result = append(result, make(DigitInstruction, 0))

		for _, letter := range line {
			parsed := Directions.Parse(letter)
			if parsed == nil {
				return nil, fmt.Errorf("parsing '%s'", string(letter))
			}

			result[lineNumber] = append(result[lineNumber], *parsed)
		}
	}

	return result, nil
}
