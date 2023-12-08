/*
For example:

LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)

Here, there are two starting nodes, 11A and 22A (because they both end with A). As you follow each left/right instruction, use that instruction to simultaneously navigate away from both nodes you're currently on. Repeat this process until all of the nodes you're currently on end with Z. (If only some of the nodes you're on end with Z, they act like any other node and you continue as normal.) In this example, you would proceed as follows:

    Step 0: You are at 11A and 22A.
    Step 1: You choose all of the left paths, leading you to 11B and 22B.
    Step 2: You choose all of the right paths, leading you to 11Z and 22C.
    Step 3: You choose all of the left paths, leading you to 11B and 22Z.
    Step 4: You choose all of the right paths, leading you to 11Z and 22B.
    Step 5: You choose all of the left paths, leading you to 11B and 22C.
    Step 6: You choose all of the right paths, leading you to 11Z and 22Z.

So, in this example, you end up entirely on nodes that end in Z after 6 steps.
*/

package main

import (
	"strings"
	"testing"

	"github.com/matryer/is"

	_ "embed"
)

//go:embed testdata/ghost_network.txt
var ghostNetwork string

func TestPart2(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   []string
		want uint64
	}{
		"empty": {[]string{}, 0},
		"simultaneously navigate away from both nodes you're currently on": {
			in:   strings.Split(ghostNetwork, "\n"),
			want: 6,
		},
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

func TestLCM(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   Ghosts
		want uint64
	}{
		"empty": {Ghosts{}, 0},
		"basic input": {
			in:   Ghosts{Ghost{stepsTaken: 2}, Ghost{stepsTaken: 3}},
			want: 6,
		},
		"wikipedia simple algorithm": {
			in:   Ghosts{Ghost{stepsTaken: 3}, Ghost{stepsTaken: 4}, Ghost{stepsTaken: 6}},
			want: 12,
		},
		"wikipedia table method": {
			in: Ghosts{
				Ghost{stepsTaken: 4},
				Ghost{stepsTaken: 7},
				Ghost{stepsTaken: 12},
				Ghost{stepsTaken: 21},
				Ghost{stepsTaken: 42},
			},
			want: 84,
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got := lcm(testCase.in)
			is.Equal(got, testCase.want)
		})
	}
}
