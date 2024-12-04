package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"

	"go.jlucktay.dev/adventofcode/2024/day02"
)

func main() {
	slog.SetDefault(
		slog.New(
			tint.NewHandler(
				os.Stderr,
				&tint.Options{TimeFormat: time.RFC3339},
			)))

	input, err := os.ReadFile("input.txt")
	if err != nil {
		slog.Error("reading file", slog.Any("err", err))
		os.Exit(1)
	}

	part1, err := day02.SafeReports(string(input))
	if err != nil {
		slog.Error("part 1", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Info("part 1", slog.Int("result", part1))

	if part1 == 0 {
		return
	}

	part2, err := day02.Part2(string(input))
	if err != nil {
		slog.Error("part 2", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Info("part 2", slog.Int("result", part2))
}
