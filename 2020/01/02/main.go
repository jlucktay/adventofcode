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
		outerInt, errConvOuter := strconv.Atoi(xInput[i])
		if errConvOuter != nil {
			continue
		}

		for j := i + 1; j < len(xInput); j++ {
			midInt, errConvMid := strconv.Atoi(xInput[j])
			if errConvMid != nil {
				continue
			}

			for k := j + 1; k < len(xInput); k++ {
				innerInt, errConvInner := strconv.Atoi(xInput[k])
				if errConvInner != nil {
					continue
				}

				if outerInt+midInt+innerInt == 2020 {
					fmt.Printf("outer '%d', mid '%d', and inner '%d' sum to 2020; product is %d\n",
						outerInt, midInt, innerInt, outerInt*midInt*innerInt)
					return
				}
			}
		}
	}
}
