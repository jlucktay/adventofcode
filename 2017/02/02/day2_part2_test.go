/*
For example, given the following spreadsheet:

5	9	2	8
9	4	7	3
3	8	6	5

In the first row, the only two numbers that evenly divide are 8 and 2; the result of this division is 4.
In the second row, the two numbers are 9 and 3; the result is 3.
In the third row, the result is 2.

In this example, the sum of the results would be 4 + 3 + 2 = 9.
*/

package main

import "testing"

func TestEvenlyDivisibleChecksum(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`0	9	2	8
9	4	0	3
3	8	6	0`, 9},
		{`5	9	2	8
9	4	7	3
3	8	6	5`, 9},
	}

	for _, c := range cases {
		got := evenlyDivisibleChecksum(c.in)

		if got != c.want {
			t.Errorf("evenlyDivisibleChecksum(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
