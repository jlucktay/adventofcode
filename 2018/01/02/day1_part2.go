// Package main for Advent of Code 2018, day 1, part 2
// http://adventofcode.com/2018/day/1
package main

func calibrateDevice(inputs []int) int {
	position := 0
	repeats := make(map[int]bool)

	for {
		for _, x := range inputs {
			if repeats[position] {
				return position
			}

			repeats[position] = true
			position += x
		}
	}
}
