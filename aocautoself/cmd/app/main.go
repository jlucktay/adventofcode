// This app will create some Advent of Code templates for you.
package main

import (
	"fmt"

	"github.com/jlucktay/adventofcode/aocautoself/pkg"
	"github.com/jlucktay/adventofcode/aocautoself/pkg/fetchaocday"
)

func main() {
	days := make([]aocautoself.Day, 0)
	days = append(days, fetchaocday.Fetch(2017, 1))
	fmt.Println(days)
}
