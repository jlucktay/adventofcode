package day02_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2024/day02"
)

func TestSafeReports(t *testing.T) {
	is := is.New(t)

	testInput := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

	result, err := day02.SafeReports(testInput)
	is.NoErr(err)
	is.Equal(2, result)
}
