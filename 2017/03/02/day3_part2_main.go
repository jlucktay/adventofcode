// Package main for Advent of Code 2017, day 3, part 2
// http://adventofcode.com/2017/day/3
package main

import "fmt"

func main() {
	var spir spiral
	spir.Init()

	target := 277678

	for index := 0; index < target; index++ {
		spir.Add()
	}

	// fmt.Println("\nlast:", spir.last)
	fmt.Println("Manhattan distance for", target, "is", spir.Manhattan())
}
