package main

import (
	"fmt"
	"os"

	aoc2022 "go.jlucktay.dev/adventofcode/2022"
	"go.jlucktay.dev/adventofcode/2022/day02"
)

func main() {
	input := aoc2022.RootCmd()

	part1, err := day02.TotalScore(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error from part 1: %v\n", err)
		os.Exit(aoc2022.EXIT_PART_1_ERROR)
	}

	fmt.Println(part1)

	part2, err := day02.StrategisedScore(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error from part 2: %v\n", err)
		os.Exit(aoc2022.EXIT_PART_2_ERROR)
	}

	fmt.Println(part2)
}
