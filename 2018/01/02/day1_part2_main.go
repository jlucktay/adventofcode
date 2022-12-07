// Package main for Advent of Code 2018, day 1, part 2
// http://adventofcode.com/2018/day/1

// What is the first frequency your device reaches twice?

package main

import (
	"fmt"

	d1 "go.jlucktay.dev/adventofcode/2018/01"
)

func main() {
	fmt.Println(calibrateDevice(d1.ProcessInput(d1.GetInput())))
}
