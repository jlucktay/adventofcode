package day13_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day13"
)

func TestPart2(t *testing.T) {
	is := is.New(t)

	actual, err := day13.Part2(INPUT)
	is.NoErr(err)
	is.Equal("Part 2", actual)
}
