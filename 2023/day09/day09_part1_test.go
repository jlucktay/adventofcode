/*
For example:

0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45

To best protect the oasis, your environmental report should include a prediction of the next value in each history. To do this, start by making a new sequence from the difference at each step of your history. If that sequence is not all zeroes, repeat this process, using the sequence you just generated as the input sequence. Once all of the values in your latest sequence are zeroes, you can extrapolate what the next value of the original history should be.

In the above dataset, the first history is 0 3 6 9 12 15. Because the values increase by 3 each step, the first sequence of differences that you generate will be 3 3 3 3 3. Note that this sequence has one fewer value than the input sequence because at each step it considers two numbers from the input. Since these values aren't all zero, repeat the process: the values differ by 0 at each step, so the next sequence is 0 0 0 0. This means you have enough information to extrapolate the history! Visually, these sequences can be arranged like this:

0   3   6   9  12  15
  3   3   3   3   3
    0   0   0   0

To extrapolate, start by adding a new zero to the end of your list of zeroes; because the zeroes represent differences between the two values above them, this also means there is now a placeholder in every sequence above it:

0   3   6   9  12  15   B
  3   3   3   3   3   A
    0   0   0   0   0

You can then start filling in placeholders from the bottom up. A needs to be the result of increasing 3 (the value to its left) by 0 (the value below it); this means A must be 3:

0   3   6   9  12  15   B
  3   3   3   3   3   3
    0   0   0   0   0

Finally, you can fill in B, which needs to be the result of increasing 15 (the value to its left) by 3 (the value below it), or 18:

0   3   6   9  12  15  18
  3   3   3   3   3   3
    0   0   0   0   0

So, the next value of the first history is 18.

Finding all-zero differences for the second history requires an additional sequence:

1   3   6  10  15  21
  2   3   4   5   6
    1   1   1   1
      0   0   0

Then, following the same process as before, work out the next value in each sequence from the bottom up:

1   3   6  10  15  21  28
  2   3   4   5   6   7
    1   1   1   1   1
      0   0   0   0

So, the next value of the second history is 28.

The third history requires even more sequences, but its next value can be found the same way:

10  13  16  21  30  45  68
   3   3   5   9  15  23
     0   2   4   6   8
       2   2   2   2
         0   0   0

So, the next value of the third history is 68.
*/

package main

import (
	"strings"
	"testing"

	"github.com/matryer/is"

	_ "embed"
)

//go:embed testdata/oasis.txt
var oasis string

func TestPart1(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   []string
		want int
	}{
		"empty":       {[]string{}, 0},
		"for example": {strings.Split(oasis, "\n"), 114},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := Part1(testCase.in)
			is.NoErr(err)
			is.Equal(got, testCase.want)
		})
	}
}

func TestReduceToZeroCount(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in      string
		want    []int
		wantLen int
	}{
		"empty":         {"", nil, 0},
		"for example 1": {"0 3 6 9 12 15", []int{3, 0}, 2},
		"for example 2": {"1 3 6 10 15 21", []int{6, 1, 0}, 3},
		"for example 3": {"10 13 16 21 30 45", []int{15, 6, 2, 0}, 4},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			hist, err := parseLine(testCase.in)
			is.NoErr(err)

			got := hist.reduceToZeroCount(false)
			is.Equal(got, testCase.want)
			is.Equal(len(got), testCase.wantLen)
		})
	}
}
