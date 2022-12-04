package day02_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day02"
)

func TestTotalScore(t *testing.T) {
	is := is.New(t)

	actual, err := day02.TotalScore(INPUT)
	is.NoErr(err)
	is.Equal(15, actual)
}
