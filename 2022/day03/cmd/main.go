package main

import (
	"fmt"

	aoc2022 "go.jlucktay.dev/adventofcode/2022"
	"go.jlucktay.dev/adventofcode/2022/day03"
)

func main() {
	input := aoc2022.RootCmd()
	fmt.Println(day03.RucksackPrioritySum(input))
	fmt.Println(day03.RucksackGroupPriority(input))
}
