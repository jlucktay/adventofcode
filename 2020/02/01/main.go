package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, errRead := os.ReadFile("../input.txt")
	if errRead != nil {
		fmt.Fprintf(os.Stderr, "could not read input: %v", errRead)
	}

	validCount, totalCount := 0, 0

	xInput := strings.Split(string(input), "\n")
	for i := range xInput {
		xLine := strings.Split(xInput[i], " ")

		xRange := strings.Split(xLine[0], "-")

		if len(xRange) < 2 {
			continue
		}

		needle := strings.Split(xLine[1], ":")[0]
		occurrences := strings.Count(xLine[2], needle)

		totalCount++

		lowerBound, _ := strconv.Atoi(xRange[0])
		upperBound, _ := strconv.Atoi(xRange[1])

		if occurrences >= lowerBound && occurrences <= upperBound {
			validCount++
		}
	}

	fmt.Printf("Valid/total count: %d/%d\n", validCount, totalCount)
}
