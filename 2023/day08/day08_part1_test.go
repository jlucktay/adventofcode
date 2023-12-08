/*
It seems like you're meant to use the left/right instructions to navigate the network. Perhaps if you have the camel follow the same instructions, you can escape the haunted wasteland!

After examining the maps for a bit, two nodes stick out: AAA and ZZZ. You feel like AAA is where you are now, and you have to follow the left/right instructions until you reach ZZZ.

This format defines each node of the network individually. For example:

RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)

Starting with AAA, you need to look up the next element based on the next left/right instruction in your input. In this example, start with AAA and go right (R) by choosing the right element of AAA, CCC. Then, L means to choose the left element of CCC, ZZZ. By following the left/right instructions, you reach ZZZ in 2 steps.
*/

package main

import (
	"cmp"
	"strings"
	"testing"

	"github.com/matryer/is"

	_ "embed"
)

//go:embed testdata/network.txt
var network string

//go:embed testdata/repeat.txt
var repeat string

func TestPart1(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   []string
		want int
	}{
		"empty":       {[]string{}, 0},
		"for example": {strings.Split(network, "\n"), 2},
		"repeat the whole sequence of instructions as necessary": {strings.Split(repeat, "\n"), 6},
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

func TestParseLines(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   string
		want Puzzle
	}{
		"empty": {"", Puzzle{}},
		"for example": {network, Puzzle{
			directions: []Direction{Right, Left},
			nodes: map[Address]Node{
				"AAA": {left: "BBB", right: "CCC"},
				"BBB": {left: "DDD", right: "EEE"},
				"CCC": {left: "ZZZ", right: "GGG"},
				"DDD": {left: "DDD", right: "DDD"},
				"EEE": {left: "EEE", right: "EEE"},
				"GGG": {left: "GGG", right: "GGG"},
				"ZZZ": {left: "ZZZ", right: "ZZZ"},
			},
		}},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := parseLines(strings.Split(testCase.in, "\n"))
			is.NoErr(err)

			for index := range got.directions {
				is.True(index < len(testCase.want.directions))
				is.Equal(0, cmp.Compare(got.directions[index], testCase.want.directions[index]))
			}

			for key, value := range got.nodes {
				_, wantExists := testCase.want.nodes[key]
				is.True(wantExists)
				is.Equal(testCase.want.nodes[key], value)
			}
		})
	}
}
