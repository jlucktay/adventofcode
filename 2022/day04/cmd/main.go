package main

import (
	"fmt"

	aoc2022 "go.jlucktay.dev/adventofcode/2022"
	"go.jlucktay.dev/adventofcode/2022/day04"
)

func main() {
	input := aoc2022.RootCmd()
	fmt.Println(day04.SectionIDFullyContain(input))
	fmt.Println(day04.SectionIDOverlap(input))
}
