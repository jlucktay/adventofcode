package day06_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day06"
)

var INPUT = []struct {
	datastream string
	marker     int
}{
	{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
	{"nppdvjthqldpwncqszvftbrmjlhg", 6},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
}

func TestFindMarkerInDatastream(t *testing.T) {
	is := is.New(t)

	for _, input := range INPUT {
		marker, err := day06.FindMarkerInDatastream(input.datastream)
		is.NoErr(err)
		is.Equal(input.marker, marker)
	}
}
