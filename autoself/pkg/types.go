// Package autoself is code that writes boilerplate code.
// It is a self-perpetuating package that gathers stories from the Advent of Code website and scaffolds out code
// skeletons for the challenge described on the given year and day.
package autoself

// Day is a single day in the Advent of Code event challenges.
// It has two DayDesc sub-structs, for parts 1 and 2 of the given day.
type Day struct {
	year, date  uint    // The year and date of this particular day
	description string  // The full text from the <article class="day-desc"> element
	dd1, dd2    DayDesc // Breakdowns of the 'description' field for the first and second parts of the day
}

// DayDesc is a breakdown of the day's challenge description into the discrete parts that feed into text templates.
type DayDesc struct {
	fluff   string // The "story" part, that introduces the problem, wrapped in some narrative
	test    string // One or more examples, to test your algorithm(s) against
	stinger string // A question that is posed and must be answered to complete the challenge
}
