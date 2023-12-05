// Package main for Advent of Code 2023, day 5, part 1
// https://adventofcode.com/2023/day/5
package main

import (
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"strings"
)

// What is the lowest location number that corresponds to any of the initial seed numbers?
func Part1(inputLines []string) (int, error) {
	seeds := []int{}
	mapsFromInput := []AtoB{}
	onAMap := false
	lowest := math.MaxInt

	for ilIndex := range inputLines {
		if inputLines[ilIndex] == "" {
			onAMap = false
			continue
		}

		if strings.HasPrefix(inputLines[ilIndex], "seeds: ") {
			inputLine := strings.TrimPrefix(inputLines[ilIndex], "seeds: ")
			xInputLine := strings.Split(inputLine, " ")
			for _, xil := range xInputLine {
				parsedXil, err := strconv.Atoi(xil)
				if err != nil {
					return 0, err
				}

				seeds = append(seeds, parsedXil)
			}

			continue
		}

		if before, found := strings.CutSuffix(inputLines[ilIndex], " map:"); found {
			onAMap = true
			xBefore := strings.Split(before, "-")
			if len(xBefore) != 3 {
				return 0, fmt.Errorf("expecting 3 tokens from line '%s' delimited by '-'", before)
			}

			newAtoB := AtoB{
				from: xBefore[0],
				to:   xBefore[2],
			}

			mapsFromInput = append(mapsFromInput, newAtoB)

			continue
		}

		if onAMap {
			xLine := strings.Split(inputLines[ilIndex], " ")
			if len(xLine) != 3 {
				return 0, fmt.Errorf("expecting 3 tokens from line '%s' delimited by '-'", inputLines[ilIndex])
			}

			parsedMapNumbers := [3]int{}

			for a, b := range xLine {
				tmp, err := strconv.Atoi(b)
				if err != nil {
					return 0, err
				}

				parsedMapNumbers[a] = tmp
			}

			mapsFromInput[len(mapsFromInput)-1].destinationRangeStart = append(
				mapsFromInput[len(mapsFromInput)-1].destinationRangeStart,
				parsedMapNumbers[0])

			mapsFromInput[len(mapsFromInput)-1].sourceRangeStart = append(
				mapsFromInput[len(mapsFromInput)-1].sourceRangeStart,
				parsedMapNumbers[1])

			mapsFromInput[len(mapsFromInput)-1].rangeLength = append(
				mapsFromInput[len(mapsFromInput)-1].rangeLength,
				parsedMapNumbers[2])
		}
	}

	slog.Info("seeds", slog.String("slice", fmt.Sprintf("%#v", seeds)))
	slog.Info("maps", slog.String("slice", fmt.Sprintf("%+v", mapsFromInput)))

	for _, seed := range seeds {
		slog.Info("outer loop",
			slog.Int("seed", seed))

		corresponding := seed

		for _, a2b := range mapsFromInput {
			slog.Info("inner loop",
				slog.Int("seed", seed),
				slog.String("from", a2b.from),
				slog.String("to", a2b.to))

			corresponding = a2b.CorrespondsTo(corresponding)
		}

		if corresponding < lowest {
			slog.Info("inside if",
				slog.Int("seed", seed),
				slog.Int("corresponding", corresponding))

			lowest = corresponding
		}
	}

	if lowest == math.MaxInt {
		lowest = 0
	}

	return lowest, nil
}

type AtoB struct {
	destinationRangeStart, sourceRangeStart, rangeLength []int
	from, to                                             string
}

func (ab AtoB) CorrespondsTo(from int) int {
	for index, srs := range ab.sourceRangeStart {
		slog.Debug("correspond",
			slog.Int("from", from),
			slog.Int("source_range_start", srs),
			slog.Int("destination_range_start", ab.destinationRangeStart[index]),
			slog.Int("range_length", ab.rangeLength[index]))

		if from >= srs {
			slog.Debug("might correspond",
				slog.Int("from", from),
				slog.Int("source_range_start", srs))

			if from < srs+ab.rangeLength[index] {
				difference := ab.destinationRangeStart[index] - srs
				answer := from + difference

				slog.Debug("should correspond",
					slog.Int("from", from),
					slog.Int("source_range_start", srs),
					slog.Int("range_length", ab.rangeLength[index]),
					slog.Int("to", answer))

				return answer
			}
		} else {
			slog.Debug("below source range start",
				slog.Int("from", from),
				slog.Int("source_range_start", srs))
		}
	}

	return from
}
