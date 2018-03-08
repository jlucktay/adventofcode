/*
http://adventofcode.com/2017/day/5
*/

package main

func followJumpOffsets(jumps []int) int {
	var curPos, nextPos, tally int

	for curPos < len(jumps) {
		nextPos += jumps[curPos]
		jumps[curPos]++
		curPos = nextPos
		tally++
	}

	return tally
}
