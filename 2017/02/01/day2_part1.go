// Package main for Advent of Code 2017, day 2, part 1
// http://adventofcode.com/2017/day/2
package main

import (
	d2 "go.jlucktay.dev/adventofcode/2017/02"
)

/*
The spreadsheet consists of rows of apparently-random numbers. To make sure the recovery process is on the right track, they need you to calculate the spreadsheet's checksum. For each row, determine the difference between the largest value and the smallest value; the checksum is the sum of all of these differences.
*/
func corruptionChecksum(input string) int {
	var totalChecksum int

	for _, row := range d2.ConvertInput(input) {
		totalChecksum += checksumSingleRow(row)
	}

	return totalChecksum
}

func checksumSingleRow(row []int) int {
	// maximum int value, which is dependent on architecture
	// from here: https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go/6878625#6878625
	lowest := int(^uint(0) >> 1)
	var highest int

	for _, cell := range row {
		if cell == 0 {
			lowest = 0
			continue
		}

		if cell > highest {
			highest = cell
		}

		if cell < lowest {
			lowest = cell
		}
	}

	return (highest - lowest)
}
