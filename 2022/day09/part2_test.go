package day09_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day09"
)

const INPUT2 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`

func TestLongTailVisitCount(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		input          string
		expectedVisits int
	}{
		"part 1": {
			input:          INPUT,
			expectedVisits: 1,
		},
		"part 2": {
			input:          INPUT2,
			expectedVisits: 36,
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase
		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			lr := day09.NewLongRope()
			is.NoErr(lr.ParseCommands(testCase.input))
			is.Equal(testCase.expectedVisits, lr.TailVisitCount())
		})
	}
}
