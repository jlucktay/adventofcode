package dayNN_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/dayNN"
)

func TestPart2(t *testing.T) {
	is := is.New(t)

	actual, err := dayNN.Part2(INPUT)
	is.NoErr(err)
	is.Equal("NN", actual)
}
