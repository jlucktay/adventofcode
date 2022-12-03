package day03_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day03"
)

func TestRucksackPrioritySum(t *testing.T) {
	is := is.New(t)

	actual, err := day03.RucksackPrioritySum(INPUT)
	is.NoErr(err)
	is.Equal(157, actual)
}
