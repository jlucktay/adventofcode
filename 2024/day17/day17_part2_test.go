/*
--- Part Two ---

Digging deeper in the device's manual, you discover the problem: this program is supposed to output another copy of the
program! Unfortunately, the value in register A seems to have been corrupted. You'll need to find a new value to which
you can initialize register A so that the program's output instructions produce an exact copy of the program itself.

For example:

Register A: 2024 Register B: 0 Register C: 0

Program: 0,3,5,4,3,0

This program outputs a copy of itself if register A is instead initialized to 117440. (The original initial value of
register A, 2024, is ignored.)

What is the lowest positive initial value for register A that causes the program to output a copy of itself?
*/

package day17_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2024/day17"
)

func TestPart2(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   string
		want int
	}{
		"empty": {"", 0},
		"for example": {`Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0
`, 117_440},
	}

	for desc, testCase := range testCases {
		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := day17.Part2(testCase.in)
			is.NoErr(err)
			is.Equal(got, testCase.want)
		})
	}
}
