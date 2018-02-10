/*
1212 produces 6: the list contains 4 items, and all four digits match the digit 2 items ahead.
1221 produces 0, because every comparison is between a 1 and a 2.
123425 produces 4, because both 2s match each other, but no other digit has a match.
123123 produces 12.
12131415 produces 4.
*/

package main

import "testing"

func TestDecode(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}

	for _, c := range cases {
		got := decode(c.in)

		if got != c.want {
			t.Errorf("decode(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
