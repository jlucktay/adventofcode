package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, errRead := ioutil.ReadFile("../input.txt")
	if errRead != nil {
		fmt.Fprintf(os.Stderr, "could not read input: %v", errRead)
	}

	xInput := strings.Split(string(input), "\n")

	for i := range xInput {
		if len(xInput[i]) == 0 {
			continue
		}

		outerInt, errConvOuter := strconv.Atoi(xInput[i])
		if errConvOuter != nil {
			fmt.Fprintf(os.Stderr, "could not convert outer input: %v", errConvOuter)
		}

		for j := i + 1; j < len(xInput); j++ {
			if len(xInput[j]) == 0 {
				continue
			}

			innerInt, errConvInner := strconv.Atoi(xInput[j])
			if errConvInner != nil {
				fmt.Fprintf(os.Stderr, "could not convert inner input: %v", errConvInner)
			}

			if outerInt+innerInt == 2020 {
				fmt.Printf("outer '%d' and inner '%d' sum to 2020; product is %d\n", outerInt, innerInt, outerInt*innerInt)
				return
			}
		}
	}
}
