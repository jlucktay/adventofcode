/*
--- Part Two ---

The Historians sure are taking a long time. To be fair, the infinite corridors are very large.

How many stones would you have after blinking a total of 75 times?
*/

package day11_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2024/day11"
)

func TestPart2(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   string
		want int
	}{
		"empty": {"", 0},
	}

	for desc, testCase := range testCases {
		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := day11.Part2(testCase.in)
			is.NoErr(err)
			is.Equal(got, testCase.want)
		})
	}
}
