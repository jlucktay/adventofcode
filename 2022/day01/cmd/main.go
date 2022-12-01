package main

import (
	"fmt"
	"os"

	"go.jlucktay.dev/adventofcode/2022/day01"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "wrong number of arguments - need an input file")
		os.Exit(1)
	}

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read file '%s': %v", os.Args[1], err)
		os.Exit(2)
	}

	fmt.Println(day01.MostCalories(string(input)))
}
