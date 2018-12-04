/*
For example, using the same list of changes above, the device would loop as follows:

Current frequency  0, change of +1; resulting frequency  1.
Current frequency  1, change of -2; resulting frequency -1.
Current frequency -1, change of +3; resulting frequency  2.
Current frequency  2, change of +1; resulting frequency  3.
(At this point, the device continues from the start of the list.)
Current frequency  3, change of +1; resulting frequency  4.
Current frequency  4, change of -2; resulting frequency  2, which has already been seen.
In this example, the first frequency reached twice is 2. Note that your device might need to repeat its list of frequency changes many times before a duplicate frequency is found, and that duplicates might be found while in the middle of processing the list.

Here are other examples:

+1, -1 first reaches 0 twice.
+3, +3, +4, -2, -4 first reaches 10 twice.
-6, +3, +8, +5, -6 first reaches 5 twice.
+7, +7, -2, -7, -4 first reaches 14 twice.
*/

package main

import "testing"

func TestCalibrateDevice(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{
			[]int{1, -2, 3, 1},
			2,
		},
		{
			[]int{1, -1},
			0,
		},
		{
			[]int{3, 3, 4, -2, -4},
			10,
		},
		{
			[]int{-6, 3, 8, 5, -6},
			5,
		},
		{
			[]int{7, 7, -2, -7, -4},
			14,
		},
	}

	for _, c := range cases {
		got := calibrateDevice(c.in)

		if got != c.want {
			t.Errorf("f(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
