package day01_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2024/day01"
)

func TestListDistance(t *testing.T) {
	is := is.New(t)

	testInput := `3   4
4   3
2   5
1   3
3   9
3   3
`

	result, err := day01.ListDistance(testInput)
	is.NoErr(err)
	is.Equal(11, result)
}
