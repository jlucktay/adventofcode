/*
http://adventofcode.com/2017/day/2
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	d2 "github.com/jlucktay/advent2017/02"
)

func main() {
	input, err := ioutil.ReadFile("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	cleanInput := strings.TrimSpace(string(input))
	fmt.Printf("Result for day 2, part 2: %d\n", evenlyDivisibleChecksum(cleanInput))
}

/*
It sounds like the goal is to find the only two numbers in each row where one evenly divides the other - that is, where the result of the division operation is a whole number. They would like you to find those numbers on each line, divide them, and add up each line's result.
*/
func evenlyDivisibleChecksum(input string) int {
	var totalChecksum int

	for _, row := range d2.ConvertInput(input) {
		cleanMatchFound := false

		for i, cellOne := range row {
			if cellOne == 0 {
				break
			}

			for _, cellTwo := range row[i+1:] {
				if cellTwo == 0 {
					break
				}

				if cellOne > cellTwo && cellOne%cellTwo == 0 {
					totalChecksum += cellOne / cellTwo
					cleanMatchFound = true
				} else if cellOne < cellTwo && cellTwo%cellOne == 0 {
					totalChecksum += cellTwo / cellOne
					cleanMatchFound = true
				}
			}

			if cleanMatchFound {
				break
			}
		}
	}

	return totalChecksum
}
