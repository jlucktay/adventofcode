// Package main for Advent of Code 2023, day 5, part 2
// https://adventofcode.com/2023/day/5
package main

import (
	"fmt"
	"log/slog"
	"math"
)

// What is the lowest location number that corresponds to any of the initial seed numbers?
func Part2(inputLines []string) (int, error) {
	seeds, mapsFromInput, err := parseInput(inputLines)
	if err != nil {
		return 0, err
	}

	slog.Debug("part 2 got parsed input")

	seedRanges := []SeedRange{}

	var nsr SeedRange
	for index, seedRange := range seeds {
		if index%2 == 0 {
			nsr = SeedRange{start: seedRange}
		} else {
			nsr.length = seedRange
			seedRanges = append(seedRanges, nsr)
		}
	}

	slog.Debug("part 2 has seed ranges ready")

	slog.Debug("seed ranges",
		slog.String("slice", fmt.Sprintf("%#v", seedRanges)))

	slog.Debug("maps",
		slog.String("slice", fmt.Sprintf("%+v", mapsFromInput)))

	lowestCandidates := make(chan int)

	for _, outerSR := range seedRanges {
		go func(sr SeedRange) {
			slog.Debug("seed range",
				slog.Int("start", sr.start),
				slog.Int("length", sr.length))

			lowest := math.MaxInt

			for seed := sr.start; seed < sr.start+sr.length; seed++ {
				corresponding := seed

				for _, a2b := range mapsFromInput {
					slog.Debug("inner loop",
						slog.Int("seed", seed),
						slog.String("from", a2b.from),
						slog.String("to", a2b.to))

					corresponding = a2b.CorrespondsTo(corresponding)
				}

				if corresponding < lowest {
					slog.Debug("inside if and corresponding is below lowest",
						slog.Int("seed", seed),
						slog.Int("corresponding", corresponding),
						slog.Int("lowest_so_far", lowest))

					lowest = corresponding
				}
			}

			lowestCandidates <- lowest
		}(outerSR)
	}

	lowestLocation := math.MaxInt

	for i := 0; i < len(seedRanges); i++ {
		tmp := <-lowestCandidates

		if tmp < lowestLocation {
			lowestLocation = tmp
		}
	}

	if lowestLocation == math.MaxInt {
		lowestLocation = 0
	}

	return lowestLocation, nil
}

type SeedRange struct {
	start, length int
}
