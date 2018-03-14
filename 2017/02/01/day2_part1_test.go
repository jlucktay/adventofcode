/*
For example, given the following spreadsheet:

5	1	9	5
7	5	3
2	4	6	8

The first row's largest and smallest values are 9 and 1, and their difference is 8.
The second row's largest and smallest values are 7 and 3, and their difference is 4.
The third row's difference is 6.

In this example, the spreadsheet's checksum would be 8 + 4 + 6 = 18.
*/

package main

import "testing"

func TestCorruptionChecksum(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`0`, 0},
		{`1	2`, 1},
		{`0	1	2`, 2},
		{`0	0	1	0	0	2	0	0`, 2},
		{`3	5	8	13	21`, 18},
		{`0	3	0	5	0	8	0	13	0	21	0`, 21},
		{`5	8	13	21	34	55`, 50},
		{`0	5	0	8	0	13	0	21	0	34	0	55`, 55},
		{`8	13	21	34	55	89	144`, 136},
		{`0	8	0	13	0	21	0	34	0	55	0	89	0	144`, 144},
		{`0	1	0
2	0	3
0	4	0	5`, 9},
		{`5	1	9	5
7	5	3
2	4	6	8`, 18},
	}

	for _, c := range cases {
		got := corruptionChecksum(c.in)

		if got != c.want {
			t.Errorf("corruptionChecksum(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
