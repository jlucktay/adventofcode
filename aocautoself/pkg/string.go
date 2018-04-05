// Package aocautoself is code that writes boilerplate code.
// It is a self-perpetuating package that gathers stories from the Advent of Code website and scaffolds out code
// skeletons for the challenge described on the given year and day.
package aocautoself

import "fmt"

func (d Day) String() string {
	s := fmt.Sprintf("Advent of Code %d, day %d.\n", d.Year, d.Date)
	s += d.Description

	return s
}
