package day02_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day02"
)

func TestStrategisedScore(t *testing.T) {
	is := is.New(t)

	actual, err := day02.StrategisedScore(INPUT)
	is.NoErr(err)
	is.Equal(12, actual)
}
