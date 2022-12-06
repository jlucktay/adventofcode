package day06_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day06"
)

var INPUT2 = []struct {
	datastream string
	marker     int
}{
	{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
	{"nppdvjthqldpwncqszvftbrmjlhg", 23},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
}

func TestFindStartOfMessageMarker(t *testing.T) {
	is := is.New(t)

	for _, input := range INPUT2 {
		marker, err := day06.FindStartOfMessageMarker(input.datastream)
		is.NoErr(err)
		is.Equal(input.marker, marker)
	}
}
