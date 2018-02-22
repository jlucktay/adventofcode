/*
For example:

Data from square 1 is carried 0 steps, since it's at the access port.
Data from square 12 is carried 3 steps, such as: down, left, left.
Data from square 23 is carried only 2 steps: up twice.
Data from square 1024 must be carried 31 steps.
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
