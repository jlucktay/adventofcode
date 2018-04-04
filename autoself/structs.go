package autoself

type Day struct {
	year, date  uint    // The year and date of this particular day
	description string  // The full text from the <article class="day-desc"> element
	dd1, dd2    DayDesc // Breakdowns of the 'description' field for the first and second parts of the day
}

type DayDesc struct {
	fluff   string // The "story" part, that introduces the problem, wrapped in some narrative
	test    string // One or more examples, to test your algorithm(s) against
	stinger string // A question that is posed and must be answered to complete the challenge
}
