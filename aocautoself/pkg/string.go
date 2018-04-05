// Package aocautoself is code that writes boilerplate code.
// It is a self-perpetuating package that gathers stories from the Advent of Code website and scaffolds out code
// skeletons for the challenge described on the given year and day.
package aocautoself

import "fmt"

func (d Day) String() string {
	s := fmt.Sprintf("Advent of Code %d, day %d.\n", d.Year, d.Date)
	s += fmt.Sprintf("Raw description: \"%s\"\n", d.Description)

	s += fmt.Sprintf("Part 1, fluff: \"%s\"\n", d.Part1.Fluff)
	s += fmt.Sprintf("Part 1, test: \"%s\"\n", d.Part1.Test)
	s += fmt.Sprintf("Part 1, stinger: \"%s\"\n", d.Part1.Stinger)

	s += fmt.Sprintf("Part 2, fluff: \"%s\"\n", d.Part2.Fluff)
	s += fmt.Sprintf("Part 2, test: \"%s\"\n", d.Part2.Test)
	s += fmt.Sprintf("Part 2, stinger: \"%s\"\n", d.Part2.Stinger)

	return s
}
