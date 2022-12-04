package aoc2022

import (
	"fmt"
	"os"
)

const (
	EXIT_SUCCESS = iota
	EXIT_NO_INPUT_FILE_ARG
	EXIT_CAN_NOT_READ_FILE
	EXIT_PART_1_ERROR
	EXIT_PART_2_ERROR
)

func RootCmd() string {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "wrong number of arguments - need an input file")
		os.Exit(EXIT_NO_INPUT_FILE_ARG)
	}

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read file '%s': %v\n", os.Args[1], err)
		os.Exit(EXIT_CAN_NOT_READ_FILE)
	}

	return string(input)
}
