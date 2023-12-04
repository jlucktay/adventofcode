// Package main for Advent of Code 2023, day 4, part 1
// https://adventofcode.com/2023/day/4
package main

import (
	"log/slog"
	"strconv"
	"strings"
)

func Part1(inputLines []string) (int, error) {
	runningTotal := 0

	for _, inputLine := range inputLines {
		if len(inputLine) == 0 {
			continue
		}

		slog.Debug("input",
			slog.String("line", inputLine))

		cardNum := strings.Split(strings.TrimSpace(inputLine), ":")
		if len(cardNum) != 2 {
			return 0, nil
		}

		xNumbers := strings.Split(strings.TrimSpace(cardNum[1]), "|")
		if len(xNumbers) != 2 {
			return 0, nil
		}

		winningNumbers := strings.Split(strings.TrimSpace(xNumbers[0]), " ")
		numbersYouHave := strings.Split(strings.TrimSpace(xNumbers[1]), " ")

		mwn := map[int]struct{}{}

		for _, wn := range winningNumbers {
			if wn == "" {
				continue
			}

			iwn, err := strconv.Atoi(wn)
			if err != nil {
				return 0, err
			}

			mwn[iwn] = struct{}{}
		}

		lineTotal := 0

		for _, nyh := range numbersYouHave {
			if nyh == "" {
				continue
			}

			inyh, err := strconv.Atoi(nyh)
			if err != nil {
				return 0, err
			}

			if _, winner := mwn[inyh]; winner {
				if lineTotal == 0 {
					lineTotal = 1
				} else {
					lineTotal *= 2
				}

				slog.Debug("winning number",
					slog.Int("number_you_have", inyh),
					slog.Int("running_total", lineTotal))
			}
		}

		runningTotal += lineTotal
	}

	return runningTotal, nil
}
