package day01

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

func TrebuchetCalibrationSum(input string) (int, error) {
	const numbers = "0123456789"

	lines := strings.Split(input, "\n")
	runningTotal := 0

	for _, line := range lines {
		firstNumberIndex := strings.IndexAny(line, numbers)
		if firstNumberIndex < 0 {
			continue
		}

		firstNumberStr := string(line[firstNumberIndex])

		lastNumberIndex := strings.LastIndexAny(line, numbers)
		if lastNumberIndex < 0 {
			continue
		}

		lastNumberStr := string(line[lastNumberIndex])

		combinedNumberStr := fmt.Sprintf("%s%s", firstNumberStr, lastNumberStr)

		combinedNumber, err := strconv.Atoi(combinedNumberStr)
		if err != nil {
			slog.Error("parsing integer",
				slog.Any("err", err),
				slog.String("input", combinedNumberStr))

			continue
		}

		runningTotal += combinedNumber
	}

	return runningTotal, nil
}
