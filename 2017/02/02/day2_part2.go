/*
http://adventofcode.com/2017/day/2
*/

package main

import (
	d2 "github.com/jlucktay/adventofcode/2017/02"
)

/*
It sounds like the goal is to find the only two numbers in each row where one evenly divides the other - that is, where the result of the division operation is a whole number. They would like you to find those numbers on each line, divide them, and add up each line's result.
*/
func evenlyDivisibleChecksum(input string) int {
	var totalChecksum int

	for _, row := range d2.ConvertInput(input) {
		totalChecksum += checksumSingleRow(row)
	}

	return totalChecksum
}

func checksumSingleRow(row []int) int {
	var rowChecksum int

	for i, cellOne := range row {
		for _, cellTwo := range row[i+1:] {
			if cellTwo != 0 && cellOne > cellTwo && cellOne%cellTwo == 0 {
				rowChecksum += cellOne / cellTwo
			} else if cellOne != 0 && cellOne < cellTwo && cellTwo%cellOne == 0 {
				rowChecksum += cellTwo / cellOne
			}
		}
	}

	return rowChecksum
}
