package day01_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2023/day01"
)

func TestTrebuchetCalibrationSum(t *testing.T) {
	is := is.New(t)

	testInput := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

	result, err := day01.TrebuchetCalibrationSum(testInput)
	is.NoErr(err)
	is.Equal(142, result)
}
