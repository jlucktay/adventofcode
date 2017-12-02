package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	corruptionChecksum()

}

func corruptionChecksum() {
	day2Input := "/Users/jameslucktaylor/Google Drive/Projects/Advent of Code 2017/day2_input.txt"
	inputArray := readLines(day2Input)

	for line := 0; line < len(inputArray); line++ {
		fmt.Println(inputArray[line])
	}
}

func readLines(filename string) []string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		//Do something
	}

	return strings.Split(string(content), "\n")
}
