package model

// NewDay initialises a new Day with the given Year (y), Date (d), empty string fields, and returns it.
func NewDay(year, date int) *Day {
	return &Day{
		Year: year,
		Date: date,

		Part1: DayDesc{},
		Part2: DayDesc{},
	}
}
