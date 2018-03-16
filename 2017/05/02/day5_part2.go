// Package main for Advent of Code 2017, day 5, part 2
// http://adventofcode.com/2017/day/5
package main

func followJumpOffsets(jumps []int) int {
	var curPos, nextPos, tally int

	for curPos < len(jumps) {
		nextPos += jumps[curPos]

		if jumps[curPos] >= 3 {
			jumps[curPos]--
		} else {
			jumps[curPos]++
		}

		curPos = nextPos
		tally++
	}

	return tally
}
