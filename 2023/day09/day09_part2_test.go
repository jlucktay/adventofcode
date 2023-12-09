/*
Of course, it would be nice to have even more history included in your report. Surely it's safe to just extrapolate backwards as well, right?

For each history, repeat the process of finding differences until the sequence of differences is entirely zero. Then, rather than adding a zero to the end and filling in the next values of each previous sequence, you should instead add a zero to the beginning of your sequence of zeroes, then fill in new first values for each previous sequence.

In particular, here is what the third example history looks like when extrapolating back in time:

5  10  13  16  21  30  45
  5   3   3   5   9  15
   -2   0   2   4   6
      2   2   2   2
        0   0   0

Adding the new values on the left side of each sequence from bottom to top eventually reveals the new left-most history value: 5.

Doing this for the remaining example data above results in previous values of -3 for the first history and 0 for the second history. Adding all three new values together produces 2.
*/

package main

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestPart2(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   []string
		want int
	}{
		"empty":       {[]string{}, 0},
		"for example": {strings.Split(oasis, "\n"), 2},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := Part2(testCase.in)
			is.NoErr(err)
			is.Equal(got, testCase.want)
		})
	}
}

func TestReduceToZeroCountBackwards(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in      string
		want    []int
		wantLen int
	}{
		"empty":         {"", nil, 0},
		"for example 1": {"0 3 6 9 12 15", []int{3, 0}, 2},
		"for example 2": {"1 3 6 10 15 21", []int{2, 1, 0}, 3},
		"for example 3": {"10 13 16 21 30 45", []int{3, 0, 2, 0}, 4},
	}
	/*
		// "for example 2": {"1 3 6 10 15 21", []int{1, 1, 0}, 3},

		0 1 3 6 10 15 21
		 1 2 3 4  5  6
		  1 1 1 1  1
		   0 0 0 0

			 "for example 3": {"10 13 16 21 30 45", []int{3, 6, 2, 0}, 4},

		 5 10 13 16 21 30  45
			5  3  3  5  9  15
			 -2 0  2  4  6
				 2 2  2  2
				  0 0  0
	*/

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			hist, err := parseLine(testCase.in)
			is.NoErr(err)

			got := hist.reduceToZeroCount(true)
			is.Equal(got, testCase.want)
			is.Equal(len(got), testCase.wantLen)
		})
	}
}

func TestPrevValue(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   string
		want int
	}{
		"empty":         {"", 0},
		"for example 1": {"0 3 6 9 12 15", -3},
		"for example 2": {"1 3 6 10 15 21", 0},
		"for example 3": {"10 13 16 21 30 45", 5},
	}
	/*
				// "for example 2": "1 3 6 10 15 21", []int{2, 1, 0}

				0   1 3 6 10 15 21
				 1   2
				  1   1
				   0   0

		 			z x
		  		 y a

					ignore 'a' and start from 'x'
					'x' minus 'y' is 'z'
					roll up from second last delta to top line

				0   1 3 6 10 15 21
				 1   2 3 4  5  6
				  1   1 1 1  1
				   0   0 0 0

				// "for example 3": "10 13 16 21 30 45", []int{3, 0, 2, 0}

				 5 10 13 16 21 30  45
					5  3  3  5  9  15
					 -2 0  2  4  6
						 2 2  2  2
						  0 0  0
	*/

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			hist, err := parseLine(testCase.in)
			is.NoErr(err)

			got := hist.prevValue()
			is.Equal(got, testCase.want)
		})
	}
}
