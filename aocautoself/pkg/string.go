// Package aocautoself is code that writes boilerplate code.
// It is a self-perpetuating package that gathers stories from the Advent of Code website and scaffolds out code
// skeletons for the challenge described on the given year and day.
package aocautoself

import "fmt"

func (d *Day) String() (s string) {
	s = fmt.Sprintf("Advent of Code %d, day %d.\n\n", d.Year, d.Date)

	s += " ---\n\n"

	s += fmt.Sprintf("Raw description: \"%s\"\n\n", d.Description)

	s += " ---\n\n"

	s += fmt.Sprintf("Part 1, fluff: \"%s\"\n\n", d.Part1.Fluff)
	s += fmt.Sprintf("Part 1, test: \"%s\"\n\n", d.Part1.Test)
	s += fmt.Sprintf("Part 1, stinger: \"%s\"\n\n", d.Part1.Stinger)

	s += " ---\n\n"

	s += fmt.Sprintf("Part 2, fluff: \"%s\"\n\n", d.Part2.Fluff)
	s += fmt.Sprintf("Part 2, test: \"%s\"\n\n", d.Part2.Test)
	s += fmt.Sprintf("Part 2, stinger: \"%s\"\n\n", d.Part2.Stinger)

	return
}
