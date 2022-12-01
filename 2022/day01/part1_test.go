package day01_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day01"
)

func TestMostCalories(t *testing.T) {
	is := is.New(t)

	actual, err := day01.MostCalories(INPUT)
	is.NoErr(err)
	is.Equal(24000, actual)
}
