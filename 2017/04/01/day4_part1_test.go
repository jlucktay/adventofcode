/*
For example:

aa bb cc dd ee is valid.
aa bb cc dd aa is not valid - the word aa appears more than once.
aa bb cc dd aaa is valid - aa and aaa count as different words.
*/

package main

import "testing"

func TestValidatePassphrase(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", false},
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	}

	for _, c := range cases {
		got := validatePassphrase(c.in)

		if got != c.want {
			t.Errorf("validatePassphrase(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
