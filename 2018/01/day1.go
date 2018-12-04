package aoc201801

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func GetInput() string {
	rawInput, err := ioutil.ReadFile("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	return string(rawInput)
}

func ProcessInput(input string) []int {
	result := make([]int, 0)

	for _, i := range strings.Split(string(input), "\n") {
		convInt, convErr := strconv.Atoi(strings.TrimSpace(string(i)))

		if convErr != nil {
			convInt = 0
		} else {
			result = append(result, convInt)
		}
	}

	return result
}
