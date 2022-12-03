package main

import (
	"fmt"

	aoc2022 "go.jlucktay.dev/adventofcode/2022"
	"go.jlucktay.dev/adventofcode/2022/day01"
)

func main() {
	input := aoc2022.RootCmd()
	fmt.Println(day01.MostCalories(input))
	fmt.Println(day01.ThreeMostCalories(input))
}
