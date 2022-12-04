package day04_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day04"
)

func TestSectionIDOverlap(t *testing.T) {
	is := is.New(t)

	actual, err := day04.SectionIDOverlap(INPUT)
	is.NoErr(err)
	is.Equal(4, actual)
}
