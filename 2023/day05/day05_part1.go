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

func parseInput(inputLines []string) ([]int, []AtoB, error) {
	seeds := []int{}
	mapsFromInput := []AtoB{}
	onAMap := false

	for ilIndex := range inputLines {
		slog.Debug("parsing input line",
			slog.String("line", inputLines[ilIndex]))

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
					return nil, nil, err
				}

				seeds = append(seeds, parsedXil)
			}

			continue
		}

		if before, found := strings.CutSuffix(inputLines[ilIndex], " map:"); found {
			onAMap = true
			xBefore := strings.Split(before, "-")
			if len(xBefore) != 3 {
				return nil, nil, fmt.Errorf("expecting 3 tokens from line '%s' delimited by '-'", before)
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
				return nil, nil, fmt.Errorf("expecting 3 tokens from line '%s' delimited by '-'", inputLines[ilIndex])
			}

			parsedMapNumbers := [3]int{}

			for a, b := range xLine {
				tmp, err := strconv.Atoi(b)
				if err != nil {
					return nil, nil, err
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

	return seeds, mapsFromInput, nil
}

// What is the lowest location number that corresponds to any of the initial seed numbers?
func Part1(inputLines []string) (int, error) {
	seeds, mapsFromInput, err := parseInput(inputLines)
	if err != nil {
		return 0, err
	}

	slog.Debug("part 1 got parsed input")

	slog.Debug("seeds",
		slog.String("slice", fmt.Sprintf("%#v", seeds)))
	slog.Debug("maps",
		slog.String("slice", fmt.Sprintf("%+v", mapsFromInput)))

	lowest := math.MaxInt

	for _, seed := range seeds {
		slog.Debug("outer loop",
			slog.Int("seed", seed))

		corresponding := seed

		for _, a2b := range mapsFromInput {
			slog.Debug("inner loop",
				slog.Int("seed", seed),
				slog.String("from", a2b.from),
				slog.String("to", a2b.to))

			corresponding = a2b.CorrespondsTo(corresponding)
		}

		if corresponding < lowest {
			slog.Debug("inside if",
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
