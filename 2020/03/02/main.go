package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	rawFile, errRead := os.ReadFile("../input.txt")
	if errRead != nil {
		fmt.Fprintf(os.Stderr, "could not read file: %v", errRead)
	}

	xLines := strings.Split(string(rawFile), "\n")

	treeCounts := make([]int, 5)

	for i, slope := range []int{1, 3, 5, 7, 1} {
		curX := 0

		for j := range xLines {
			fmt.Printf("[%03d] %s", j, xLines[j])

			if j == 0 || len(xLines[j]) == 0 {
				fmt.Println()
				continue
			}

			if i == 4 && j%2 != 0 {
				fmt.Println()
				continue
			}

			curX += slope
			curX %= len(xLines[j])

			if string(xLines[j][curX]) == "#" {
				treeCounts[i]++
			}

			fmt.Printf("\t[%2d] '%s'\n", curX, string(xLines[j][curX]))
		}

		fmt.Println()
	}

	fmt.Printf("Trees hit: %#v, ", treeCounts)

	product := treeCounts[0]

	for a := range treeCounts {
		if a == 0 {
			continue
		}

		product *= treeCounts[a]
	}

	fmt.Printf("product: %d\n", product)
}
