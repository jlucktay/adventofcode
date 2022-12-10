package main

import (
	"fmt"
	"os"

	aoc2022 "go.jlucktay.dev/adventofcode/2022"
	"go.jlucktay.dev/adventofcode/2022/day10"
)

func main() {
	input := aoc2022.RootCmd()

	part1, err := day10.SumSixSignalStrengths(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error from part 1: %v\n", err)
		os.Exit(aoc2022.EXIT_PART_1_ERROR)
	}

	fmt.Println(part1)

	part2, err := day10.RenderImage(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error from part 2: %v\n", err)
		os.Exit(aoc2022.EXIT_PART_2_ERROR)
	}

	fmt.Println(part2)
}
