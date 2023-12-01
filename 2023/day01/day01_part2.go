package day01

import "strings"

func TrebuchetCalibrationSumPartTwo(input string) (int, error) {
	numbersAsStrings := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	lines := strings.Split(input, "\n")
	runningTotal := 0

	for _, line := range lines {
		stringIndex := 0
		updatedLine := strings.Builder{}

		for stringIndex < len(line) {
			lineFragment := line[stringIndex:]

			foundSomething := false

			for numberString := range numbersAsStrings {
				if strings.HasPrefix(lineFragment, numberString) {
					updatedLine.WriteString(numbersAsStrings[numberString])

					foundSomething = true
					break
				}
			}

			if !foundSomething {
				updatedLine.WriteByte(lineFragment[0])
			}

			stringIndex++
		}

		updated := updatedLine.String()

		lineSum, err := TrebuchetCalibrationSum(updated)
		if err != nil {
			return 0, err
		}

		runningTotal += lineSum
	}

	return runningTotal, nil
}
