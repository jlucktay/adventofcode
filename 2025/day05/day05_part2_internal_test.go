package day05

import (
	"testing"

	"github.com/matryer/is"
)

func TestOverlapsWith(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		range1, range2 freshRange
		want           bool
	}{
		"one": {
			range1: freshRange{lowerBound: 3, upperBound: 5},
			range2: freshRange{lowerBound: 10, upperBound: 14},
			want:   false,
		},
		"two": {
			range1: freshRange{lowerBound: 10, upperBound: 14},
			range2: freshRange{lowerBound: 16, upperBound: 20},
			want:   false,
		},
		"three": {
			range1: freshRange{lowerBound: 16, upperBound: 20},
			range2: freshRange{lowerBound: 12, upperBound: 18},
			want:   true,
		},
		"four": {
			range1: freshRange{lowerBound: 12, upperBound: 18},
			range2: freshRange{lowerBound: 16, upperBound: 20},
			want:   true,
		},
		"five": {
			range1: freshRange{lowerBound: 12, upperBound: 20},
			range2: freshRange{lowerBound: 16, upperBound: 18},
			want:   true,
		},
		"six": {
			range1: freshRange{lowerBound: 16, upperBound: 18},
			range2: freshRange{lowerBound: 12, upperBound: 20},
			want:   true,
		},
		"seven": {
			range1: freshRange{lowerBound: 1, upperBound: 4},
			range2: freshRange{lowerBound: 2, upperBound: 6},
			want:   true,
		},
	}

	for desc, testCase := range testCases {
		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got := testCase.range1.overlapsWith(testCase.range2)
			is.Equal(got, testCase.want)
		})
	}
}
