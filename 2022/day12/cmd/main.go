package main

import (
	"fmt"
	"os"

	aoc2022 "go.jlucktay.dev/adventofcode/2022"
	"go.jlucktay.dev/adventofcode/2022/day12"
)

func main() {
	input := aoc2022.RootCmd()

	part1, err := day12.Part1(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error from part 1: %v\n", err)
		os.Exit(aoc2022.EXIT_PART_1_ERROR)
	}

	fmt.Println(part1)

	part2, err := day12.Part2(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error from part 2: %v\n", err)
		os.Exit(aoc2022.EXIT_PART_2_ERROR)
	}

	fmt.Println(part2)
}
