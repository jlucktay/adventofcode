package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	testInput := [][]int{{5, 1, 9, 5}, {7, 5, 3}, {2, 4, 6, 8}}
	fmt.Println(corruptionChecksum(testInput), "should match 18")

	day2Input := "/Users/jameslucktaylor/Google Drive/Projects/Advent of Code 2017/day2_input.txt"
	inputArray := readLines(day2Input)
	fmt.Println(corruptionChecksum(inputArray))
}

func corruptionChecksum(input [][]int) int {
	totalChecksum := 0

	for _, rowContent := range input {
		lowest := 9999
		highest := 0

		for _, cell := range rowContent {
			switch {
			case cell > highest:
				highest = cell
			case cell < lowest:
				lowest = cell
			}
		}

		totalChecksum += (highest - lowest)
	}

	return totalChecksum
}

func evenlyDivisibleChecksum(input []string) int {
	totalChecksum := 0

	for row := 0; row < len(input); row++ {
		cells := strings.Split(input[row], "\t")
		lowest := 9999
		highest := 0

		for cell := 0; cell < len(cells); cell++ {
			intCell, _ := strconv.Atoi(cells[cell])

			if intCell > highest {
				highest = intCell
			}

			if intCell < lowest {
				lowest = intCell
			}
		}

		totalChecksum += (highest - lowest)
	}

	return totalChecksum
}

func readLines(filename string) [][]int {
	var arrReturn = [][]int{}
	content, ioErr := ioutil.ReadFile(filename)

	if ioErr != nil {
		panic(ioErr)
	}

	for row, i := range strings.Split(string(content), "\n") {
		for _, j := range strings.Split(i, "\t") {
			convInt, convErr := strconv.Atoi(string(j))

			if convErr != nil {
				panic(convErr)
			}

			arrReturn[row] = append(arrReturn[row], convInt)
		}
	}

	return arrReturn
}

//	Failed to continue:
//		Cannot find Delve debugger.
//		Install from https://github.com/derekparker/delve & ensure it is in your "GOPATH/bin" or "PATH"."
