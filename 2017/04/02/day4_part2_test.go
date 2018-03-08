/*
For example:

abcde fghij is a valid passphrase.
abcde xyz ecdab is not valid - the letters from the third word can be rearranged to form the first word.
a ab abc abd abf abj is a valid passphrase, because all letters need to be used when forming another word.
iiii oiii ooii oooi oooo is valid.
oiii ioii iioi iiio is not valid - any of these words can be rearranged to form any other word.
*/

package main

import "testing"

func TestValidatePassphrase(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", false},
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false},
	}

	for _, c := range cases {
		got := validatePassphrase(c.in)

		if got != c.want {
			t.Errorf("validatePassphrase(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
