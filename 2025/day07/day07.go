// Package day07 for Advent of Code 2025, day 7.
// https://adventofcode.com/2025/day/7
package day07

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"log/slog"
	"slices"
	"strings"

	"github.com/orsinium-labs/enum"
)

type Tile enum.Member[rune]

var (
	Empty    = Tile{'.'}
	Splitter = Tile{'^'}
	Start    = Tile{'S'}
	Beam     = Tile{'|'}

	Tiles = enum.New(Empty, Splitter, Start, Beam)
)

type TachyonBeamSplitter struct {
	grid       map[image.Point]Tile
	timelines  map[image.Point]int
	xMax, yMax int
}

func (tbs TachyonBeamSplitter) String() string {
	sb := strings.Builder{}

	for y := range tbs.yMax {
		for x := range tbs.xMax {
			currentPos := image.Pt(x, y)

			if timelineCount, exists := tbs.timelines[currentPos]; exists {
				if timelineCount >= 10 {
					sb.WriteRune('x')
				} else {
					sb.WriteString(fmt.Sprintf("%d", timelineCount))
				}
			} else {
				sb.WriteRune(tbs.grid[currentPos].Value)
			}
		}

		sb.WriteString("\n")
	}

	return sb.String()
}

func (tbs TachyonBeamSplitter) startSplitting() int {
	result := 0

	for y := 1; y < tbs.yMax; y++ {
		for x := range tbs.xMax {
			currentPosition := image.Pt(x, y)
			positionAbove := currentPosition.Sub(image.Pt(0, 1))

			if tbs.grid[positionAbove] == Start || tbs.grid[positionAbove] == Beam {
				switch tbs.grid[currentPosition] {
				case Empty:
					tbs.grid[currentPosition] = Beam

				case Splitter:
					leftOfCurrent := currentPosition.Sub(image.Pt(1, 0))
					rightOfCurrent := currentPosition.Add(image.Pt(1, 0))

					leftSplit, rightSplit := false, false

					if tbs.grid[leftOfCurrent] == Empty {
						tbs.grid[leftOfCurrent] = Beam
						leftSplit = true

						slog.Debug("left of splitter was empty", slog.Any("current", currentPosition), slog.Any("left", leftOfCurrent), slog.Int("result", result))
					}

					if tbs.grid[rightOfCurrent] == Empty {
						tbs.grid[rightOfCurrent] = Beam
						rightSplit = true

						slog.Debug("right of splitter was empty", slog.Any("current", currentPosition), slog.Any("right", rightOfCurrent), slog.Int("result", result))
					}

					if leftSplit || rightSplit {
						result++
					}
				}
			}
		}
	}

	return result
}

func (tbs TachyonBeamSplitter) startQuantum() int {
	result := 0

CreatedFirstBlock:
	for y := 1; y < tbs.yMax; y++ {
		for x := range tbs.xMax {
			currentPosition := image.Pt(x, y)
			positionAbove := currentPosition.Sub(image.Pt(0, 1))

			if tbs.grid[positionAbove] == Start {
				// Create the first '1' block
				tbs.timelines[currentPosition] = 1

				break CreatedFirstBlock
			}
		}
	}

	for y := 1; y < tbs.yMax; y++ {
		for x := range tbs.xMax {
			currentPosition := image.Pt(x, y)

			if currentTimelineValue, exists := tbs.timelines[currentPosition]; exists {
				below := currentPosition.Add(image.Pt(0, 1))

				switch tbs.grid[below] {
				case Empty:
					slog.Debug("nothing below current position", slog.Any("currentPosition", currentPosition), slog.Int("currentTimelineValue", currentTimelineValue))

					tbs.timelines[below] += currentTimelineValue

				case Splitter:
					slog.Debug("splitter is below current position", slog.Any("currentPosition", currentPosition), slog.Int("currentTimelineValue", currentTimelineValue))

					belowLeft := currentPosition.Add(image.Pt(-1, 1))
					belowRight := currentPosition.Add(image.Pt(1, 1))

					tbs.timelines[belowLeft] += currentTimelineValue
					tbs.timelines[belowRight] += currentTimelineValue
				}
			}
		}

		if y == tbs.yMax-1 {
			slog.Debug("bottom row; add up the blocks")

			for x := range tbs.xMax {
				if bottomRowBlockValue, exists := tbs.timelines[image.Pt(x, y)]; exists {
					result += bottomRowBlockValue
				}
			}
		}
	}

	return result
}

func parseInput(input string) (TachyonBeamSplitter, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := TachyonBeamSplitter{
		grid:      make(map[image.Point]Tile),
		timelines: make(map[image.Point]int),
	}

	y := 0

	for scanner.Scan() {
		xLine := strings.Split(scanner.Text(), " ")
		xLine = slices.DeleteFunc(xLine, func(s string) bool { return s == "" })

		if len(xLine) != 1 {
			return TachyonBeamSplitter{}, fmt.Errorf("unexpected line format '%+v'", xLine)
		}

		x := 0

		for _, r := range xLine[0] {
			newTile := Tiles.Parse(r)
			if newTile == nil {
				return TachyonBeamSplitter{}, fmt.Errorf("unknown tile '%s'", string(r))
			}

			newPoint := image.Pt(x, y)
			result.grid[newPoint] = *newTile

			x++
		}

		if result.xMax < x {
			result.xMax = x
		}

		y++
	}

	result.yMax = y

	return result, nil
}
