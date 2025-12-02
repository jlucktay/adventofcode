// Package day01 for Advent of Code 2025, day 1.
// https://adventofcode.com/2025/day/1
package day01

import (
	"bufio"
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Safe struct {
	dial                 int
	leftPointingAtZero   int
	anyClickPointsAtZero int
}

func NewSafe() *Safe {
	return &Safe{dial: 50}
}

func (s *Safe) DialPointsToZero() bool {
	return s.dial == 0
}

type Rotation struct {
	left      bool
	magnitude int
}

func (s *Safe) Turn(input Rotation) {
	for range input.magnitude {
		if input.left {
			s.dial--
		} else {
			s.dial++
		}

		if s.dial < 0 {
			s.dial = 99
		}

		if s.dial > 99 {
			s.dial = 0
		}

		if s.DialPointsToZero() {
			s.anyClickPointsAtZero++
		}
	}
}

type Instructions []Rotation

func (s *Safe) Follow(input Instructions) {
	for _, inst := range input {
		s.Turn(inst)

		if s.DialPointsToZero() {
			s.leftPointingAtZero++
		}
	}
}

func parseInput(input string) (Instructions, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := make(Instructions, 0)

	for scanner.Scan() {
		xLine := strings.Split(scanner.Text(), " ")
		xLine = slices.DeleteFunc(xLine, func(s string) bool { return s == "" })

		if len(xLine) != 1 {
			return nil, fmt.Errorf("unexpected line format '%+v'", xLine)
		}

		newInstruction := Rotation{}

		if strings.HasPrefix(xLine[0], "L") {
			newInstruction.left = true
		} else if strings.HasPrefix(xLine[0], "R") {
			newInstruction.left = false
		} else {
			return nil, fmt.Errorf("unexpected prefix on line '%s'", xLine[0])
		}

		newMagnitude, err := strconv.Atoi(xLine[0][1:])
		if err != nil {
			return nil, fmt.Errorf("converting magnitude on line '%s': %w", xLine[0], err)
		}

		newInstruction.magnitude = newMagnitude

		result = append(result, newInstruction)
	}

	return result, nil
}
