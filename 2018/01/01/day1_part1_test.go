/*
For example, if the device displays frequency changes of +1, -2, +3, +1, then starting from a frequency of zero, the following changes would occur:

Current frequency  0, change of +1; resulting frequency  1.
Current frequency  1, change of -2; resulting frequency -1.
Current frequency -1, change of +3; resulting frequency  2.
Current frequency  2, change of +1; resulting frequency  3.
In this example, the resulting frequency is 3.

Here are other example situations:

+1, +1, +1 results in  3
+1, +1, -2 results in  0
-1, -2, -3 results in -6
*/

package main

import (
	"testing"
)

func TestResultingFrequency(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{
			[]int{1, -2, 3, 1},
			3,
		},
		{
			[]int{1, 1, 1},
			3,
		},
		{
			[]int{1, 1, -2},
			0,
		},
		{
			[]int{-1, -2, -3},
			-6,
		},
	}

	for _, c := range cases {
		got := resultingFrequency(c.in)

		if got != c.want {
			t.Errorf("f(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
