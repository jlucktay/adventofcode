package day04

import (
	"testing"

	"github.com/matryer/is"
)

func TestSearchWord(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   string
		want int
	}{
		"empty": {"", 0},
		"GfG": {`abab
abeb
ebeb
`, 3},
	}

	for desc, testCase := range testCases {
		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			ws, err := parseInput(testCase.in)
			is.NoErr(err)

			is.Equal(searchWord(ws, "abe"), testCase.want)
		})
	}
}
