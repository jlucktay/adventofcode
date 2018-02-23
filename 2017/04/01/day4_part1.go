/*
http://adventofcode.com/2017/day/4
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	var tally uint

	for _, p := range strings.Split(string(input[:]), "\n") {
		if validatePassphrase(p) {
			tally++
		}
	}

	fmt.Printf("Result for day 4, part 1: %d\n", tally)
}

func validatePassphrase(passphrase string) bool {
	if len(passphrase) == 0 {
		return false
	}

	split := strings.Split(passphrase, " ")

	for a, b := range split {
		for _, d := range split[a+1:] {
			if b == d {
				return false
			}
		}
	}

	return true
}
