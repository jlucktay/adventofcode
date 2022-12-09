package day09_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day09"
)

func TestTailVisitCount(t *testing.T) {
	is := is.New(t)

	r := day09.NewRope()
	is.NoErr(r.ParseCommands(INPUT))
	is.Equal(13, r.TailVisitCount())
}
