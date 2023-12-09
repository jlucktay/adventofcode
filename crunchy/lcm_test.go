package crunchy_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/crunchy"
)

func TestLCM(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   []int
		want int
	}{
		"empty": {
			in:   []int{},
			want: 0,
		},
		"basic input": {
			in:   []int{2, 3},
			want: 6,
		},
		"wikipedia simple algorithm": {
			in:   []int{3, 4, 6},
			want: 12,
		},
		"wikipedia table method": {
			in:   []int{4, 7, 12, 21, 42},
			want: 84,
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got := crunchy.LCM(testCase.in...)
			is.Equal(got, testCase.want)
		})
	}
}

// 0637042177
