// Package day13 for Advent of Code 2024, day 13.
// https://adventofcode.com/2024/day/13
package day13

import (
	"fmt"
	"image"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	a, b, p image.Point
}

type Machines []Machine

func parseInput(input string) (Machines, error) {
	xRawMachines := strings.Split(input, "\n\n")

	result := make(Machines, 0)

	for _, rawMachine := range xRawMachines {
		if len(rawMachine) == 0 {
			continue
		}

		xRawMachine := strings.Split(rawMachine, "\n")

		rawButtonA := xRawMachine[0]
		rawButtonB := xRawMachine[1]
		rawPrize := xRawMachine[2]

		rawMachineFields := func(r rune) bool { return r == ' ' || r == ':' || r == '+' || r == ',' || r == '=' }

		xRawButtonA := strings.FieldsFunc(rawButtonA, rawMachineFields)
		xRawButtonB := strings.FieldsFunc(rawButtonB, rawMachineFields)
		xRawPrize := strings.FieldsFunc(rawPrize, rawMachineFields)

		rawMachineDelete := func(s string) bool { return s == "" }

		xRawButtonA = slices.DeleteFunc(xRawButtonA, rawMachineDelete)
		xRawButtonB = slices.DeleteFunc(xRawButtonB, rawMachineDelete)
		xRawPrize = slices.DeleteFunc(xRawPrize, rawMachineDelete)

		buttonAX, err := strconv.Atoi(xRawButtonA[3])
		if err != nil {
			return nil, fmt.Errorf("parsing: %w", err)
		}

		buttonAY, err := strconv.Atoi(xRawButtonA[5])
		if err != nil {
			return nil, fmt.Errorf("parsing: %w", err)
		}

		buttonBX, err := strconv.Atoi(xRawButtonB[3])
		if err != nil {
			return nil, fmt.Errorf("parsing: %w", err)
		}

		buttonBY, err := strconv.Atoi(xRawButtonB[5])
		if err != nil {
			return nil, fmt.Errorf("parsing: %w", err)
		}

		prizeX, err := strconv.Atoi(xRawPrize[2])
		if err != nil {
			return nil, fmt.Errorf("parsing: %w", err)
		}

		prizeY, err := strconv.Atoi(xRawPrize[4])
		if err != nil {
			return nil, fmt.Errorf("parsing: %w", err)
		}

		newMachine := Machine{
			a: image.Point{X: buttonAX, Y: buttonAY},
			b: image.Point{X: buttonBX, Y: buttonBY},
			p: image.Point{X: prizeX, Y: prizeY},
		}

		result = append(result, newMachine)
	}

	return result, nil
}
