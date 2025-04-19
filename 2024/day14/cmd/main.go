package main

import (
	"image"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"

	"go.jlucktay.dev/adventofcode/2024/day14"
)

func main() {
	slog.SetDefault(
		slog.New(
			tint.NewHandler(
				os.Stderr,
				&tint.Options{TimeFormat: time.RFC3339},
			)))

	bounds := image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{101, 103},
	}
	part1Duration := 100

	startPart1 := time.Now()
	part1, err := day14.Part1(input, bounds, part1Duration)
	finishPart1 := time.Since(startPart1)

	if err != nil {
		slog.Error("part 1", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Info("part 1", slog.Int("result", part1), slog.Duration("elapsed", finishPart1))

	if part1 == 0 {
		return
	}

	part2Duration := 1_000_000

	startPart2 := time.Now()
	part2, err := day14.Part2(input, bounds, part2Duration)
	finishPart2 := time.Since(startPart2)

	if err != nil {
		slog.Error("part 2", slog.Any("err", err))
		os.Exit(2)
	}

	slog.Info("part 2", slog.Int("result", part2), slog.Duration("elapsed", finishPart2))
}
