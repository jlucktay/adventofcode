/*
--- Part Two ---

During the bathroom break, someone notices that these robots seem awfully similar to ones built and used at the North Pole. If they're the same type of robots, they should have a hard-coded Easter egg: very rarely, most of the robots should arrange themselves into a picture of a Christmas tree.

What is the fewest number of seconds that must elapse for the robots to display the Easter egg?
*/

package day14_test

import (
	"image"
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2024/day14"
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

			got, err := day14.Part2(testCase.in, image.Rectangle{}, 0)
			is.NoErr(err)
			is.Equal(got, testCase.want)
		})
	}
}
