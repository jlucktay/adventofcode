/*
So, the first few squares' values are chosen as follows:

Square 1 starts with the value 1.
Square 2 has only one adjacent filled square (with value 1), so it also stores 1.
Square 3 has both of the above squares as neighbors and stores the sum of their values, 2.
Square 4 has all three of the aforementioned squares as neighbors and stores the sum of their values, 4.
Square 5 only has the first and fourth squares as neighbors, so it gets the value 5.
*/

package main

import "testing"

func TestLastValue(t *testing.T) {
	cases := []struct {
		nodesToAdd uint
		lastValue  uint
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 4},
		{5, 5},
		{6, 10},
		{7, 11},
		{8, 23},
		{9, 25},
		{10, 26},
		{11, 54},
		{12, 57},
		{13, 59},
		{14, 122},
		{15, 133},
		{16, 142},
		{17, 147},
		{18, 304},
		{19, 330},
		{20, 351},
		{21, 362},
		{22, 747},
		{23, 806},
	}

	for _, c := range cases {
		var s spiral
		s.Init()

		for i := uint(0); i < c.nodesToAdd; i++ {
			s.Add()
		}

		got := s.last.value

		if got != c.lastValue {
			t.Errorf("Value of last node after adding %v nodes == %v, wanted %v", c.nodesToAdd, got, c.lastValue)
		}
	}
}

func TestManhattan(t *testing.T) {
	cases := []struct {
		in   int
		want uint64
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}

	for _, c := range cases {
		var s spiral
		s.Init()

		for i := 0; i < c.in; i++ {
			s.Add()
		}

		got := s.Manhattan()

		if got != c.want {
			t.Errorf("Manhattan(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestCoordsString(t *testing.T) {
	cases := []struct {
		in   spiralNodeCoords
		want string
	}{
		{spiralNodeCoords{0, 0}, "[0,0]"},
		{spiralNodeCoords{-100, -1000}, "[-100,-1000]"},
		{spiralNodeCoords{123, 456}, "[123,456]"},
	}

	for _, c := range cases {
		if c.in.String() != c.want {
			t.Errorf("String of '%v' doesn't match '%v'", c.in.String(), c.want)
		}
	}
}
