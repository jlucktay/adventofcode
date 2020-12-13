package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	rawInput, errRead := ioutil.ReadFile("../input.txt")
	if errRead != nil {
		fmt.Fprintf(os.Stderr, "could not read file: %v", errRead)
	}

	xPassports := strings.Split(string(rawInput), "\n\n")
	validPassports, totalPassports := 0, 0

	for i := range xPassports {
		xPassports[i] = strings.ReplaceAll(xPassports[i], "\n", " ")
		xPassports[i] = strings.TrimSpace(xPassports[i])

		xFields := strings.Fields(xPassports[i])
		fieldMap := map[string]string{}

		for j := range xFields {
			xSplit := strings.Split(xFields[j], ":")
			fieldMap[xSplit[0]] = xSplit[1]
		}

		fmt.Printf("[%3d] %v\n", i, fieldMap)

		totalPassports++

		if _, hasCID := fieldMap["cid"]; hasCID {
			if len(fieldMap) >= 8 {
				validPassports++
			}
		} else if len(fieldMap) >= 7 {
			validPassports++
		}
	}

	fmt.Printf("Valid/total: %d/%d\n", validPassports, totalPassports)
}
