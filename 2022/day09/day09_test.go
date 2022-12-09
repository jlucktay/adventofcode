package day09_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day09"
)

const INPUT = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

func TestDistanceFrom(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		head, tail       day09.Position
		expectedDistance int
	}{
		"initial state": {
			head:             day09.NewPosition(0, 0),
			tail:             day09.NewPosition(0, 0),
			expectedDistance: 0,
		},
		"after R 4": {
			head:             day09.NewPosition(4, 0),
			tail:             day09.NewPosition(0, 0),
			expectedDistance: 4,
		},
		"after U 4": {
			head:             day09.NewPosition(4, 4),
			tail:             day09.NewPosition(0, 0),
			expectedDistance: 4,
		},
		"after L 3": {
			head:             day09.NewPosition(1, 4),
			tail:             day09.NewPosition(0, 0),
			expectedDistance: 4,
		},
		"after D 1": {
			head:             day09.NewPosition(1, 3),
			tail:             day09.NewPosition(0, 0),
			expectedDistance: 3,
		},
		"after R 4 again": {
			head:             day09.NewPosition(5, 3),
			tail:             day09.NewPosition(0, 0),
			expectedDistance: 5,
		},
		"after D 1 again": {
			head:             day09.NewPosition(5, 2),
			tail:             day09.NewPosition(0, 0),
			expectedDistance: 5,
		},
		"after L 5": {
			head:             day09.NewPosition(0, 2),
			tail:             day09.NewPosition(0, 0),
			expectedDistance: 2,
		},
		"after R 2": {
			head:             day09.NewPosition(2, 2),
			tail:             day09.NewPosition(0, 0),
			expectedDistance: 2,
		},
		"diagonal one": {
			head:             day09.NewPosition(1, 1),
			tail:             day09.NewPosition(0, 0),
			expectedDistance: 1,
		},
		"diagonal two": {
			head:             day09.NewPosition(2, 2),
			tail:             day09.NewPosition(0, 0),
			expectedDistance: 2,
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase
		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			is.Equal(testCase.expectedDistance, testCase.tail.DistanceFrom(testCase.head))
		})
	}
}
