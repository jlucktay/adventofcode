package day10

// Find the signal strength during the 20th, 60th, 100th, 140th, 180th, and 220th cycles.
// What is the sum of these six signal strengths?

func SumSixSignalStrengths(input string) (int, error) {
	cc := NewClockCircuit()

	if err := cc.ParseProgram(input); err != nil {
		return 0, err
	}

	tape := cc.Tape()
	str := tape[19]*20 + tape[59]*60 + tape[99]*100 + tape[139]*140 + tape[179]*180 + tape[219]*220

	return str, nil
}
