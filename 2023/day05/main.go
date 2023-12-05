// Package main for Advent of Code 2023, day 5
// https://adventofcode.com/2023/day/5
package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

const inputPath = "input.txt"

func main() {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		slog.Error("could not read input file",
			slog.Any("err", err),
			slog.String("path", inputPath))

		return
	}

	lines := strings.Split(string(input), "\n")

	p1, err := Part1(lines)
	if err != nil {
		slog.Error("part 1",
			slog.Any("err", err))

		return
	}

	fmt.Printf("Part 1: %d\n", p1)

	p2, err := Part2(lines)
	if err != nil {
		slog.Error("part 2",
			slog.Any("err", err))

		return
	}

	fmt.Printf("Part 2: %d\n", p2)
}
