package main

import (
	"fmt"

	aoc2022 "go.jlucktay.dev/adventofcode/2022"
	"go.jlucktay.dev/adventofcode/2022/day02"
)

func main() {
	input := aoc2022.RootCmd()
	fmt.Println(day02.TotalScore(input))
	fmt.Println(day02.StrategisedScore(input))
}
