package main

import (
	"fmt"
	"os"

	aoc2022 "go.jlucktay.dev/adventofcode/2022"
	day13 "go.jlucktay.dev/adventofcode/2022/day13"
)

func main() {
	input := aoc2022.RootCmd()

	part1, err := day13.Part1(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error from part 1: %v\n", err)
		os.Exit(aoc2022.EXIT_PART_1_ERROR)
	}

	fmt.Println(part1)

	part2, err := day13.Part2(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error from part 2: %v\n", err)
		os.Exit(aoc2022.EXIT_PART_2_ERROR)
	}

	fmt.Println(part2)
}
