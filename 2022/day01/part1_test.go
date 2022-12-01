package day01_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day01"
)

func TestMostCalories(t *testing.T) {
	is := is.New(t)

	const INPUT string = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

	actual, err := day01.MostCalories(INPUT)
	is.NoErr(err)
	is.Equal(24000, actual)
}
