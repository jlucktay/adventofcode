package day08_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day08"
)

func TestTreesVisibleFromOutsideGrid(t *testing.T) {
	is := is.New(t)

	actual, err := day08.TreesVisibleFromOutsideGrid(day08.INPUT)
	is.NoErr(err)
	is.Equal(21, actual)
}
