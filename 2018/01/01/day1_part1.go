// Package main for Advent of Code 2018, day 1, part 1
// http://adventofcode.com/2018/day/1
package main

func resultingFrequency(inputs []int) int {
	result := 0

	for _, x := range inputs {
		result += x
	}

	return result
}
