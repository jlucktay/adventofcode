/*
http://adventofcode.com/2017/day/4
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
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

	fmt.Printf("Result for day 4, part 2: %d\n", tally)
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)

	return strings.Join(s, "")
}

func validatePassphrase(passphrase string) bool {
	split := strings.Split(passphrase, " ")

	if len(passphrase) == 0 {
		return false
	}

	for a, b := range split {
		outer := sortString(b)

		for _, d := range split[a+1:] {
			inner := sortString(d)

			if outer == inner {
				return false
			}
		}
	}

	return true
}
