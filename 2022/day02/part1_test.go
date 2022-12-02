package day02_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day02"
)

func TestTotalScore(t *testing.T) {
	is := is.New(t)

	is.Equal(day02.TotalScore(INPUT), 15)
}
