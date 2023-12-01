package day01_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2023/day01"
)

func TestTrebuchetCalibrationSumPartTwo(t *testing.T) {
	is := is.New(t)

	testInput := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

	result, err := day01.TrebuchetCalibrationSumPartTwo(testInput)
	is.NoErr(err)
	is.Equal(281, result)
}

func TestTrebuchetCalibrationSumPartTwoPerLine(t *testing.T) {
	is := is.New(t)

	testLines := map[string]int{
		"two1nine":         29,
		"eightwothree":     83,
		"abcone2threexyz":  13,
		"xtwone3four":      24,
		"4nineeightseven2": 42,
		"zoneight234":      14,
		"7pqrstsixteen":    76,

		"oneight": 18,
	}

	for a, b := range testLines {
		input, expected := a, b

		t.Run(input, func(t *testing.T) {
			is := is.New(t)

			actual, err := day01.TrebuchetCalibrationSumPartTwo(input)
			is.NoErr(err)
			is.Equal(expected, actual)
		})
	}
}
