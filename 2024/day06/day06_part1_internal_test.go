package day06

import (
	"testing"

	"github.com/matryer/is"
)

func TestString(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   string
		want string
	}{
		"empty": {"", ""},
		"for example": {`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`, `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`},
	}

	for desc, testCase := range testCases {
		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			m, err := parseInput(testCase.in)
			is.NoErr(err)

			got := m.String()
			is.Equal(got, testCase.want)
		})
	}
}
