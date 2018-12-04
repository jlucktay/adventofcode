// Package main for Advent of Code 2018, day 1, part 1
// http://adventofcode.com/2018/day/1

// Starting with a frequency of zero, what is the resulting frequency after all of the changes in frequency have been applied?

package main

import (
	"fmt"

	d1 "github.com/jlucktay/adventofcode/2018/01"
)

func main() {
	fmt.Println(resultingFrequency(d1.ProcessInput(d1.GetInput())))
}
