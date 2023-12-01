package main

import (
	"fmt"
	"os"

	"go.jlucktay.dev/adventofcode/2023/day01"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	fmt.Println(day01.TrebuchetCalibrationSum(string(input)))
}
