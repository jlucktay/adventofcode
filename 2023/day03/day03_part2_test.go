/*
Consider the same engine schematic again:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..

In this schematic, there are two gears.
The first is in the top left; it has part numbers 467 and 35, so its gear ratio is 16345.
The second gear is in the lower right; its gear ratio is 451490.
(The * adjacent to 617 is not a gear because it is only adjacent to one part number.)
Adding up all of the gear ratios produces 467835.
*/

package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestPart2(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   []string
		want int
	}{
		"empty": {[]string{}, 0},
		"Consider the same engine schematic again": {
			[]string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			467835,
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

func TestLookAroundGear(t *testing.T) {
	is := is.New(t)

	schematicRunes := [][]rune{
		{'4', '6', '7', '.', '.', '1', '1', '4', '.', '.'},
		{'.', '.', '.', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '3', '5', '.', '.', '6', '3', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
		{'6', '1', '7', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '+', '.', '5', '8', '.'},
		{'.', '.', '5', '9', '2', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '7', '5', '5', '.'},
		{'.', '.', '.', '$', '.', '*', '.', '.', '.', '.'},
		{'.', '6', '6', '4', '.', '5', '9', '8', '.', '.'},
	}

	testCases := map[string]struct {
		in          [][]rune
		row, column int
		want        [2]int
	}{
		"empty": {[][]rune{}, 0, 0, [2]int{0, 0}},
		"first - second line": {
			schematicRunes,
			1, 3,
			[2]int{467, 35},
		},
		"second - fifth line - not actually a gear": {
			schematicRunes,
			4, 3,
			[2]int{0, 0},
		},
		"third - ninth line": {
			schematicRunes,
			8, 5,
			[2]int{755, 598},
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := lookAroundGear(testCase.in, testCase.row, testCase.column)
			is.NoErr(err)
			is.Equal(got, testCase.want)
		})
	}
}

func TestDiscoverEntireNumber(t *testing.T) {
	is := is.New(t)

	schematicRunes := [][]rune{
		{'4', '6', '7', '.', '.', '1', '1', '4', '.', '.'},
		{'.', '.', '.', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '3', '5', '.', '.', '6', '3', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
		{'6', '1', '7', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '+', '.', '5', '8', '.'},
		{'.', '.', '5', '9', '2', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '7', '5', '5', '.'},
		{'.', '.', '.', '$', '.', '*', '.', '.', '.', '.'},
		{'.', '6', '6', '4', '.', '5', '9', '8', '.', '.'},
	}

	testCases := map[string]struct {
		in          [][]rune
		row, column int
		want        int
	}{
		"empty": {[][]rune{}, 0, 0, 0},
		"first - 467": {
			schematicRunes,
			0, 0,
			467,
		},
		"second - 114": {
			schematicRunes,
			0, 5,
			114,
		},
		"third - 35": {
			schematicRunes,
			2, 2,
			35,
		},
		"third - 35 from the 5": {
			schematicRunes,
			2, 3,
			35,
		},
		"fourth - 633 0": {
			schematicRunes,
			2, 6,
			633,
		},
		"fourth - 633 1": {
			schematicRunes,
			2, 7,
			633,
		},
		"fourth - 633 2": {
			schematicRunes,
			2, 8,
			633,
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := discoverEntireNumber(testCase.in, testCase.row, testCase.column)
			is.NoErr(err)
			is.Equal(got, testCase.want)
		})
	}
}
