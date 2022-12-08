package day08_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day08"
)

func TestHighestScenicScore(t *testing.T) {
	is := is.New(t)

	actual, err := day08.HighestScenicScore(day08.INPUT)
	is.NoErr(err)
	is.Equal(8, actual)
}
