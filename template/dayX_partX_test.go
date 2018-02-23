/*
For example...
*/

package main

import "testing"

func TestFunction(t *testing.T) {
	cases := []struct {
		in   bool
		want bool
	}{
		{true, true},
	}

	for _, c := range cases {
		got := f(c.in)

		if got != c.want {
			t.Errorf("f(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
