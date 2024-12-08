// Package day08 for Advent of Code 2024, day 8.
// https://adventofcode.com/2024/day/8
package day08

import (
	"bufio"
	"bytes"
	"fmt"
	"slices"
	"strings"

	"github.com/orsinium-labs/enum"
)

type GridPointThing enum.Member[rune]

var (
	gptBuilder = enum.NewBuilder[rune, GridPointThing]()

	Antinode = gptBuilder.Add(GridPointThing{'#'})
	Empty    = gptBuilder.Add(GridPointThing{'.'})

	GridPointThings = gptBuilder.Enum()
)

type GridPoint []GridPointThing

type AntennaGrid struct {
	theGrid        [][]GridPoint
	thingDirectory map[GridPointThing][][2]int
}

func (ag AntennaGrid) String() string {
	sb := strings.Builder{}

	for _, row := range ag.theGrid {
		for _, gp := range row {
			if len(gp) > 1 {
				if len(gp) == 2 && slices.Contains(gp, Empty) && slices.Contains(gp, Antinode) {
					sb.WriteRune(Antinode.Value)
				} else {
					sb.WriteRune('%')
				}
			} else if len(gp) == 1 {
				sb.WriteRune(gp[0].Value)
			} else {
				panic("we're on a GridPoint without any things")
			}
		}

		sb.WriteString("\n")
	}

	sb.WriteString("\n")

	for gpt, locations := range ag.thingDirectory {
		if gpt == Empty {
			continue
		}

		gptLocationReport := fmt.Sprintf("%s:", string(gpt.Value))

		for _, coordinates := range locations {
			gptLocationReport += fmt.Sprintf(" %d,%d", coordinates[0], coordinates[1])
		}

		sb.WriteString(gptLocationReport + "\n")
	}

	sb.WriteString("\nGridPointThings:")

	for _, gptv := range GridPointThings.Values() {
		sb.WriteString(fmt.Sprintf(" '%s'", string(gptv)))
	}

	sb.WriteString("\n")

	return sb.String()
}

func parseInput(input string) (AntennaGrid, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := AntennaGrid{
		theGrid:        make([][]GridPoint, 0),
		thingDirectory: make(map[GridPointThing][][2]int),
	}

	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		result.theGrid = append(result.theGrid, make([]GridPoint, 0))

		for x, r := range line {
			newGPT := GridPointThing{r}

			if !GridPointThings.Contains(newGPT) {
				gptBuilder.Add(newGPT)
				GridPointThings = gptBuilder.Enum()
			}

			result.theGrid[y] = append(result.theGrid[y], GridPoint{newGPT})

			if _, exists := result.thingDirectory[newGPT]; !exists {
				result.thingDirectory[newGPT] = make([][2]int, 0)
			}

			result.thingDirectory[newGPT] = append(result.thingDirectory[newGPT], [2]int{x, y})
		}

		y++
	}

	for y, gridRow := range result.theGrid {
		for x, gridPoints := range gridRow {
			if slices.Contains(gridPoints, Antinode) {
				panic(fmt.Sprintf("there is an antinode on a freshly-parsed grid at x: %d, y: %d", x, y))
			}

			if len(gridPoints) != 1 {
				panic(fmt.Sprintf(
					"there is a grid point on a freshly-parsed grid at x: %d, y: %d with more/less than one thing on it: %#v",
					x, y, gridPoints))
			}
		}
	}

	return result, nil
}
