package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	rawFile, errRead := ioutil.ReadFile("../input.txt")
	if errRead != nil {
		fmt.Fprintf(os.Stderr, "could not read file: %v", errRead)
	}

	xLines := strings.Split(string(rawFile), "\n")
	curX, treeCount := 0, 0

	for i := range xLines {
		fmt.Printf("[%03d] %s", i, xLines[i])

		if i == 0 || len(xLines[i]) == 0 {
			fmt.Println()
			continue
		}

		curX += 3
		curX %= len(xLines[i])

		if string(xLines[i][curX]) == "#" {
			treeCount++
		}

		fmt.Printf("\t[%2d] '%s'\n", curX, string(xLines[i][curX]))
	}

	fmt.Printf("Trees hit: %d\n", treeCount)
}
