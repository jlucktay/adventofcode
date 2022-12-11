package model

// Day is a single day in the Advent of Code event challenges.
// It has two DayDesc sub-structs, for parts 1 and 2 of the given day.
type Day struct {
	// The year and date of this particular day
	Year, Date int

	// The raw text from the <article class="day-desc"> element(s)
	Description string

	// Breakdowns of the 'description' field for the first and second parts of the day
	Part1, Part2 DayDesc
}

// DayDesc is a breakdown of the day's challenge description into the discrete parts that feed into text templates.
type DayDesc struct {
	// The "story" part, that introduces the problem, wrapped in some narrative
	Fluff string

	// One or more examples, to test your algorithm(s) against
	Test string

	// A question that is posed and must be answered to complete the challenge
	Stinger string
}
