package day06_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day06"
)

func TestFindStartOfMessageMarker(t *testing.T) {
	is := is.New(t)

	for _, input := range INPUT {
		marker, err := day06.FindStartOfMessageMarker(input.datastream)
		is.NoErr(err)
		is.Equal(input.message, marker)
	}
}
