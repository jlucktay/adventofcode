package dayNN_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/dayNN"
)

const INPUT = ``

func TestPart1(t *testing.T) {
	is := is.New(t)

	actual, err := dayNN.Part1(INPUT)
	is.NoErr(err)
	is.Equal("NN", actual)
}
