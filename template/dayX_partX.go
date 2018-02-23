/*
http://adventofcode.com/.../day/...
*/

package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	_ = strings.Split(strings.TrimSpace(string(input[:])), "\n")
}

func f(b bool) bool {
	return b
}
