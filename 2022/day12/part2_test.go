package day12_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day12"
)

func TestPart2(t *testing.T) {
	is := is.New(t)

	actual, err := day12.Part2(INPUT)
	is.NoErr(err)
	is.Equal(29, actual)
}
