/*
For example...
*/

package day19_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2024/day19"
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

			got, err := day19.Part2(testCase.in)
			is.NoErr(err)
			is.Equal(got, testCase.want)
		})
	}
}
