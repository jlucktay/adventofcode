package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestHistoryString(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   string
		want string
	}{
		"empty": {"", ""},
		"for example": {
			"190 10 19 3267 81 40 27 83 17 5 156 15 6 7290 6 8 6 15 161011 16 10 13 192 17 8 14 21037 9 7 18 13 292 11 6",
			"190 10 19 3267 81 40 27 83 17 5 156 15 6 7290 6 8 6 15 161011 16 10 13 192 17 8 14 21037 9 7 18 13 292 11 6",
		},
	}

	for desc, testCase := range testCases {
		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := parseLine(testCase.in)
			is.NoErr(err)
			is.Equal(got.String(), testCase.want)
		})
	}
}
