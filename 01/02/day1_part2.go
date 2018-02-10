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

	fmt.Printf("Result for day 1, part 2: %d\n", decode(cleanInput))
}

/*
Now, instead of considering the next digit, it wants you to consider the digit halfway around the circular list. That is, if your list contains 10 items, only include a digit in your sum if the digit 10/2 = 5 steps forward matches it. Fortunately, your list has an even number of elements.
*/
func decode(encoded string) uint64 {
	halfLength := len(encoded) / 2
	var sum uint64

	for i, c := range encoded {
		if i >= halfLength {
			sum *= 2
			break
		}

		current, _ := strconv.ParseUint(string(c), 10, 8)
		halfwayForward := encoded[(i+halfLength)%len(encoded)]
		halfway, _ := strconv.ParseUint(string(halfwayForward), 10, 8)

		if current == halfway {
			sum += current
		}
	}

	return sum
}
