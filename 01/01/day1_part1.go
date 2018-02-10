/*
http://adventofcode.com/2017/day/1
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	cleanInput := strings.TrimSpace(string(input))
	fmt.Printf("Result for day 1, part 1: %d\n", decode(cleanInput))
}

/*
The captcha requires you to review a sequence of digits (your puzzle input) and find the sum of all digits that match the next digit in the list. The list is circular, so the digit after the last digit is the first digit in the list.
*/
func decode(encoded string) int {
	var current, previous, sum int

	for _, c := range encoded {
		previous = current
		current, _ = strconv.Atoi(string(c))

		if previous == current {
			sum += current
		}
	}

	if encoded[0] == encoded[len(encoded)-1] {
		finalMatch, _ := strconv.Atoi(string(encoded[0]))
		sum += finalMatch
	}

	return sum
}
