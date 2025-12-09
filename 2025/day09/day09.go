// Package day09 for Advent of Code 2025, day 9.
// https://adventofcode.com/2025/day/9
package day09

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"log/slog"
	"slices"
	"strconv"
	"strings"
)

type MovieTheater struct {
	corners []image.Point
}

func (mt *MovieTheater) largestRectangle() int {
	result := 0

	for i := 0; i < len(mt.corners)-1; i++ {
		for j := i + 1; j < len(mt.corners); j++ {
			slog.Debug("considering candidate rectangle", slog.Any("firstCorner", mt.corners[i]), slog.Any("secondCorner", mt.corners[j]))

			size := image.Rectangle{
				Min: mt.corners[i],
				Max: mt.corners[j],
			}.Size()

			slog.Debug("size of raw rectangle", slog.Any("size", size))

			// Flip negatives.
			if size.X < 0 {
				size.X *= -1
			}

			if size.Y < 0 {
				size.Y *= -1
			}

			// Increment to include the extra area.
			size.X++
			size.Y++

			slog.Debug("size of modified candidate rectangle", slog.Any("size", size))

			if area := size.X * size.Y; area > result {
				slog.Debug("found a bigger rectangle", slog.Any("firstCorner", mt.corners[i]), slog.Any("secondCorner", mt.corners[j]), slog.Int("area", area))
				result = area
			}
		}
	}

	return result
}

func parseInput(input string) (MovieTheater, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := MovieTheater{
		corners: make([]image.Point, 0),
	}

	for scanner.Scan() {
		xLine := strings.Split(scanner.Text(), ",")
		xLine = slices.DeleteFunc(xLine, func(s string) bool { return s == "" })

		if len(xLine) != 2 {
			return MovieTheater{}, fmt.Errorf("unexpected line format '%+v'", xLine)
		}

		ptX, err := strconv.Atoi(xLine[0])
		if err != nil {
			return MovieTheater{}, fmt.Errorf("converting X coordinate '%s': %w", xLine[0], err)
		}

		ptY, err := strconv.Atoi(xLine[1])
		if err != nil {
			return MovieTheater{}, fmt.Errorf("converting Y coordinate '%s': %w", xLine[1], err)
		}

		result.corners = append(result.corners, image.Pt(ptX, ptY))
	}

	return result, nil
}
