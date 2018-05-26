// Package aocday holds functions for creating and manipulating the Day type.
package aocday

import (
	"github.com/jlucktay/adventofcode/aocautoself/pkg"
)

// NewDay initialises a new Day with the given Year (y), Date (d), empty string fields, and returns it.
func NewDay(y, d uint) (day *aocautoself.Day) {
	day = &aocautoself.Day{}

	day.Year = y
	day.Date = d
	day.Description = ""

	day.Part1 = aocautoself.DayDesc{Fluff: "", Test: "", Stinger: ""}
	day.Part2 = aocautoself.DayDesc{Fluff: "", Test: "", Stinger: ""}

	return
}
