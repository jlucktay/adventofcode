package main

import (
	"fmt"
	"os"

	"go.jlucktay.dev/adventofcode/2022/day03"
)

const (
	EXIT_SUCCESS = iota
	EXIT_NO_INPUT_FILE_ARG
	EXIT_CAN_NOT_READ_FILE
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "wrong number of arguments - need an input file")
		os.Exit(EXIT_NO_INPUT_FILE_ARG)
	}

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read file '%s': %v\n", os.Args[1], err)
		os.Exit(EXIT_CAN_NOT_READ_FILE)
	}

	fmt.Println(day03.RucksackPrioritySum(string(input)))
	fmt.Println(day03.RucksackGroupPriority(string(input)))
}
