// Package main for Advent of Code 2017, day 1, part 2
// http://adventofcode.com/2017/day/1
package main

import "strconv"

/*
Now, instead of considering the next digit, it wants you to consider the digit halfway around the circular list. That is, if your list contains 10 items, only include a digit in your sum if the digit 10/2 = 5 steps forward matches it. Fortunately, your list has an even number of elements.
*/
func decode(encoded string) int {
	halfLength := len(encoded) / 2
	var sum int

	for i, c := range encoded {
		if i >= halfLength {
			sum *= 2
			break
		}

		current, _ := strconv.Atoi(string(c))
		halfwayForward := encoded[(i+halfLength)%len(encoded)]
		halfway, _ := strconv.Atoi(string(halfwayForward))

		if current == halfway {
			sum += current
		}
	}

	return sum
}
