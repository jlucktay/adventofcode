package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, errRead := ioutil.ReadFile("../input.txt")
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

		totalCount++

		firstPos, _ := strconv.Atoi(xRange[0])
		secondPos, _ := strconv.Atoi(xRange[1])

		// Convert to indices
		firstPos--
		secondPos--

		haystack := xLine[2]

		firstHay := string([]rune(haystack)[firstPos])
		secondHay := string([]rune(haystack)[secondPos])

		if (firstHay == needle) != (secondHay == needle) {
			validCount++
		}
	}

	fmt.Printf("Valid/total count: %d/%d\n", validCount, totalCount)
}
