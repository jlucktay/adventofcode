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

func TestManhattan(t *testing.T) {
	cases := []struct {
		in   int
		want uint
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

// Co-ordinates always have 8 neighbours surrounding them
func TestNeighboursCount(t *testing.T) {
	cases := []struct {
		in   spiralNodeCoords
		want int
	}{
		{
			spiralNodeCoords{0, 0},
			8,
		},
		{
			spiralNodeCoords{27, 42},
			8,
		},
		{
			spiralNodeCoords{-100, -200},
			8,
		},
	}

	for _, c := range cases {
		if len(c.in.neighbours()) != 8 {
			t.Errorf("len(%v.neighbours()) != 8", c.in)
		}
	}
}

func TestCoordsString(t *testing.T) {
	cases := []struct {
		in   spiralNodeCoords
		want string
	}{
		{
			spiralNodeCoords{0, 0},
			"[0,0]",
		},
		{
			spiralNodeCoords{-100, -1000},
			"[-100,-1000]",
		},
		{
			spiralNodeCoords{123, 456},
			"[123,456]",
		},
	}

	for _, c := range cases {
		if c.in.String() != c.want {
			t.Errorf("String of '%v' doesn't match '%v'", c.in.String(), c.want)
		}
	}
}
