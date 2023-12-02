package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"

	"go.jlucktay.dev/adventofcode/2023/day02"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	slog.SetDefault(
		slog.New(
			tint.NewHandler(
				os.Stderr,
				&tint.Options{
					TimeFormat: time.RFC3339,
				})))

	part1, err := day02.CubeConundrum(string(input))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	slog.Info("part 1", slog.Int("result", part1))

	part2, err := day02.CubeConundrumPartTwo(string(input))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	slog.Info("part 2", slog.Int("result", part2))
}
