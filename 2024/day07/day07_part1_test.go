/*
For example:

190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20

Each line represents a single equation. The test value appears before the colon on each line; it is your job to
determine whether the remaining numbers can be combined with operators to produce the test value.

Operators are always evaluated left-to-right, not according to precedence rules. Furthermore, numbers in the equations
cannot be rearranged. Glancing into the jungle, you can see elephants holding two different types of operators: add (+)
and multiply (*).

Only three of the above equations can be made true by inserting operators:

- 190: 10 19 has only one position that accepts an operator: between 10 and 19. Choosing + would give 29, but choosing
  * would give the test value (10 * 19 = 190).
- 3267: 81 40 27 has two positions for operators. Of the four possible configurations of the operators, two cause the
  right side to match the test value: 81 + 40 * 27 and 81 * 40 + 27 both equal 3267 (when evaluated left-to-right)!
- 292: 11 6 16 20 can be solved in exactly one way: 11 + 6 * 16 + 20.

The engineers just need the total calibration result, which is the sum of the test values from just the equations that
could possibly be true. In the above example, the sum of the test values for the three equations listed above is 3749.

Determine which equations could possibly be true. What is their total calibration result?
*/

package day07_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2024/day07"
)

func TestPart1(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   string
		want uint64
	}{
		"empty": {"", 0},
		"for example": {`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`, 3749},
	}

	for desc, testCase := range testCases {
		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := day07.Part1(testCase.in)
			is.NoErr(err)
			is.Equal(got, testCase.want)
		})
	}
}
