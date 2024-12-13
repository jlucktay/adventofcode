// Package day13 for Advent of Code 2024, day 13, part 1.
// https://adventofcode.com/2024/day/13
package day13

func Part1(input string) (int, error) {
	machines, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := 0

	for _, machine := range machines {
		result += calculate(machine)
	}

	return result, nil
}

func calculate(m Machine) int {
	// The coordinates of the target (prize) are some number of A button presses and B button presses away.

	// Calculate necessary number of presses for each button.
	pressA := (m.b.Y*m.p.X - m.b.X*m.p.Y) / (m.a.X*m.b.Y - m.a.Y*m.b.X)
	pressB := (m.a.Y*m.p.X - m.a.X*m.p.Y) / (m.a.Y*m.b.X - m.a.X*m.b.Y)

	// If the number of presses for both buttons is bang on target, return the cost.
	if m.a.Mul(pressA).Add(m.b.Mul(pressB)) == m.p {
		return pressA*3 + pressB
	}

	// However, if we're off the mark at all, return a zero to indicate that it's not possible to land exactly.
	return 0
}
