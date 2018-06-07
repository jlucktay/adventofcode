// Package main for Advent of Code 2017, day 3, part 2
// http://adventofcode.com/2017/day/3
package main

import "fmt"

func main() {
	var spir spiral
	spir.Init()
	spir.Add()

	target := uint(277678)

	for spir.last.value < target {
		spir.Add()
	}

	fmt.Println("last:", spir.last)
}
