// Package main for Advent of Code 2017, day 4, part 1
// http://adventofcode.com/2017/day/4
package main

import (
	"strings"
)

func validatePassphrase(passphrase string) bool {
	if len(passphrase) == 0 {
		return false
	}

	split := strings.Split(passphrase, " ")

	for a, b := range split {
		for _, d := range split[a+1:] {
			if b == d {
				return false
			}
		}
	}

	return true
}
