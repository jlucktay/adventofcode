package model

import "fmt"

func (d *Day) String() string {
	s := fmt.Sprintf("Advent of Code %d, day %d.\n\n", d.Year, d.Date)

	if d.Description != "" {
		s += " ---\n\n"
		s += fmt.Sprintf("Raw description: \"%s\"\n\n", d.Description)
	}

	if d.Part1.Fluff != "" || d.Part1.Stinger != "" || d.Part1.Test != "" {
		s += " ---\n\n"
		s += fmt.Sprintf("Part 1, fluff: \"%s\"\n\n", d.Part1.Fluff)
		s += fmt.Sprintf("Part 1, test: \"%s\"\n\n", d.Part1.Test)
		s += fmt.Sprintf("Part 1, stinger: \"%s\"\n\n", d.Part1.Stinger)
	}

	if d.Part2.Fluff != "" || d.Part2.Stinger != "" || d.Part2.Test != "" {
		s += " ---\n\n"
		s += fmt.Sprintf("Part 2, fluff: \"%s\"\n\n", d.Part2.Fluff)
		s += fmt.Sprintf("Part 2, test: \"%s\"\n\n", d.Part2.Test)
		s += fmt.Sprintf("Part 2, stinger: \"%s\"\n\n", d.Part2.Stinger)
	}

	return s
}
