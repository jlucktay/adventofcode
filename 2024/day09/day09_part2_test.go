/*
--- Part Two ---

Upon completion, two things immediately become clear. First, the disk definitely has a lot more contiguous free space,
just like the amphipod hoped. Second, the computer is running much more slowly! Maybe introducing all of that file
system fragmentation was a bad idea?

The eager amphipod already has a new plan: rather than move individual blocks, he'd like to try compacting the files on
his disk by moving whole files instead.

This time, attempt to move whole files to the leftmost span of free space blocks that could fit the file. Attempt to
move each file exactly once in order of decreasing file ID number starting with the file with the highest file ID
number. If there is no span of free space to the left of a file that is large enough to fit the file, the file does not
move.

The first example from above now proceeds differently:

00...111...2...333.44.5555.6666.777.888899
0099.111...2...333.44.5555.6666.777.8888..
0099.1117772...333.44.5555.6666.....8888..
0099.111777244.333....5555.6666.....8888..
00992111777.44.333....5555.6666.....8888..

The process of updating the filesystem checksum is the same; now, this example's checksum would be 2858.

Start over, now compacting the amphipod's hard drive using this new method instead. What is the resulting filesystem
checksum?
*/

package day09_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2024/day09"
)

func TestPart2(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   string
		want int64
	}{
		"empty":       {"", 0},
		"12345":       {"12345", 132},
		"for example": {"2333133121414131402", 2858},
	}

	/*
		12345

		0..111....22222

		0 * 0 = 0
		1 * 0 = 0
		2 * 0 = 0
		3 * 1 = 3
		4 * 1 = 4
		5 * 1 = 5
		6 * 0 = 0
		7 * 0 = 0
		8 * 0 = 0
		9 * 0 = 0
		10 * 2 = 20
		11 * 2 = 22
		12 * 2 = 24
		13 * 2 = 26
		14 * 2 = 28

		132
	*/

	for desc, testCase := range testCases {
		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := day09.Part2(testCase.in)
			is.NoErr(err)
			is.Equal(got, testCase.want)
		})
	}
}
