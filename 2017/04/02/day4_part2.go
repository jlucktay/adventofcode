// Package main for Advent of Code 2017, day 4, part 2
// http://adventofcode.com/2017/day/4
package main

import (
	"sort"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)

	return strings.Join(s, "")
}

func validatePassphrase(passphrase string) bool {
	split := strings.Split(passphrase, " ")

	if len(passphrase) == 0 {
		return false
	}

	for a, b := range split {
		outer := sortString(b)

		for _, d := range split[a+1:] {
			inner := sortString(d)

			if outer == inner {
				return false
			}
		}
	}

	return true
}
