package day11_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day11"
)

func TestTwentyRoundsOfMonkeyBusiness(t *testing.T) {
	is := is.New(t)

	actual, err := day11.TwentyRoundsOfMonkeyBusiness(INPUT)
	is.NoErr(err)
	is.Equal(10_605, actual)
}
