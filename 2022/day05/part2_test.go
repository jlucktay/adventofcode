package day05_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day05"
)

func TestTopCrate9001(t *testing.T) {
	is := is.New(t)

	actual, err := day05.TopCrate9001(INPUT)
	is.NoErr(err)
	is.Equal("MCD", actual)
}
