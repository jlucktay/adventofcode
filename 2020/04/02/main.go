package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	// BirthYear has four digits; at least 1920 and at most 2002.
	BirthYear = "byr"

	// IssueYear has four digits; at least 2010 and at most 2020.
	IssueYear = "iyr"

	// ExpirationYear has four digits; at least 2020 and at most 2030.
	ExpirationYear = "eyr"

	// Height is a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	Height = "hgt"

	// HairColor is a '#' followed by exactly six characters 0-9 or a-f.
	HairColor = "hcl"

	// EyeColor is exactly one of: amb blu brn gry grn hzl oth.
	EyeColor = "ecl"

	// PassportID is a nine-digit number, including leading zeroes.
	PassportID = "pid"

	// CountryID is ignored, missing or not.
	CountryID = "cid"
)

func main() {
	rawInput, errRead := os.ReadFile("../input.txt")
	if errRead != nil {
		fmt.Fprintf(os.Stderr, "could not read file: %v\n", errRead)
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

			switch xSplit[0] {
			case BirthYear:
			case IssueYear:
			case ExpirationYear:
			case Height:
			case HairColor:
			case EyeColor:
			case PassportID:
			case CountryID:
			default:
				fmt.Fprintf(os.Stderr, "unknown passport field: %s\n", xSplit[0])
			}

			fieldMap[xSplit[0]] = xSplit[1]
		}

		fmt.Printf("[%3d] %v\n", i, fieldMap)

		totalPassports++

		if cidValue, hasCountryID := fieldMap[CountryID]; hasCountryID && len(cidValue) > 0 {
			if len(fieldMap) >= 8 {
				validPassports++
			}
		} else if len(fieldMap) >= 7 {
			validPassports++
		}
	}

	fmt.Printf("Valid/total: %d/%d\n", validPassports, totalPassports)
}
