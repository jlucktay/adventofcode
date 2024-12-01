package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"

	"go.jlucktay.dev/adventofcode/2024/day01"
)

func main() {
	slog.SetDefault(
		slog.New(
			tint.NewHandler(
				os.Stderr,
				&tint.Options{TimeFormat: time.RFC3339},
			)))

	const inputTxt = "input.txt"

	input, err := os.ReadFile(inputTxt)
	if err != nil {
		slog.Error("reading file", slog.Any("err", err))
		os.Exit(1)
	}

	part1, err := day01.ListDistance(string(input))
	if err != nil {
		slog.Error("part 1", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Info("part 1", slog.Int("result", part1))
}
