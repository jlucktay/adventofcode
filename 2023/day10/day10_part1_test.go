/*
For example, here is a square loop of pipe:

.....
.F-7.
.|.|.
.L-J.
.....

If the animal had entered this loop in the northwest corner, the sketch would instead look like this:

.....
.S-7.
.|.|.
.L-J.
.....

In the above diagram, the S tile is still a 90-degree F bend: you can tell because of how the adjacent pipes connect to it.

Unfortunately, there are also many pipes that aren't connected to the loop! This sketch shows the same loop as above:

-L|F7
7S-7|
L|7||
-L-J|
L|-JF

In the above diagram, you can still figure out which pipes form the main loop: they're the ones connected to S, pipes those pipes connect to, pipes those pipes connect to, and so on. Every pipe in the main loop connects to its two neighbors (including S, which will have exactly two pipes connecting to it, and which is assumed to connect back to those two pipes).

Here is a sketch that contains a slightly more complex main loop:

..F7.
.FJ|.
SJ.L7
|F--J
LJ...

Here's the same example sketch with the extra, non-main-loop pipe tiles also shown:

7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ

If you want to get out ahead of the animal, you should find the tile in the loop that is farthest from the starting position. Because the animal is in the pipe, it doesn't make sense to measure this by direct distance. Instead, you need to find the tile that would take the longest number of steps along the loop to reach from the starting point - regardless of which way around the loop the animal went.

In the first example with the square loop:

.....
.S-7.
.|.|.
.L-J.
.....

You can count the distance each tile in the loop is from the starting point like this:

.....
.012.
.1.3.
.234.
.....

In this example, the farthest point from the start is 4 steps away.

Here's the more complex loop again:

..F7.
.FJ|.
SJ.L7
|F--J
LJ...

Here are the distances for each tile on that loop:

..45.
.236.
01.78
14567
23...
*/

package main

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

var squareLoop = `.....
.S-7.
.|.|.
.L-J.
.....
`

var squareLoopWithUnconnected = `-L|F7
7S-7|
L|7||
-L-J|
L|-JF
`

var slightlyMoreComplexMainLoop = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`

var withExtraNonMainLoopPipeTiles = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ
`

func TestPart1(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   []string
		want int
	}{
		"empty":                               {[]string{}, 0},
		"square loop of pipe":                 {strings.Split(squareLoop, "\n"), 4},
		"square loop with unconnected":        {strings.Split(squareLoopWithUnconnected, "\n"), 4},
		"slightly more complex main loop":     {strings.Split(slightlyMoreComplexMainLoop, "\n"), 8},
		"with extra non-main-loop pipe tiles": {strings.Split(withExtraNonMainLoopPipeTiles, "\n"), 8},
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
