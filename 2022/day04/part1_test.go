package day04_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day04"
)

func TestSectionIDFullyContain(t *testing.T) {
	is := is.New(t)

	actual, err := day04.SectionIDFullyContain(INPUT)
	is.NoErr(err)
	is.Equal(2, actual)
}
