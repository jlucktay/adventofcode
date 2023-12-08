// Package main for Advent of Code 2023, day 8, part 1
// https://adventofcode.com/2023/day/8
package main

import (
	"fmt"
	"strings"
)

type Direction uint8

const (
	Left Direction = iota
	Right
)

type Address string

type Puzzle struct {
	directions []Direction
	nodes      map[Address]Node
}

type Node struct {
	left, right Address
}

func parseLines(inputLines []string) (Puzzle, error) {
	puzzle := Puzzle{
		directions: make([]Direction, 0),
		nodes:      make(map[Address]Node),
	}

	for ilIndex := range inputLines {
		if ilIndex == 0 {
			var err error

			puzzle.directions, err = parseDirections(inputLines[ilIndex])
			if err != nil {
				return Puzzle{}, err
			}

			continue
		}

		if len(inputLines[ilIndex]) == 0 {
			continue
		}

		newAddress, newNode, err := parseNode(inputLines[ilIndex])
		if err != nil {
			return Puzzle{}, err
		}

		puzzle.nodes[Address(newAddress)] = newNode
	}

	return puzzle, nil
}

func parseDirections(input string) ([]Direction, error) {
	directions := make([]Direction, 0)

	for _, r := range input {
		switch r {
		case 'L':
			directions = append(directions, Left)

		case 'R':
			directions = append(directions, Right)

		default:
			return []Direction{}, fmt.Errorf("rune '%s' is not 'L' nor 'R'", string(r))
		}
	}

	return directions, nil
}

func parseNode(input string) (string, Node, error) {
	node := Node{}

	xInput := strings.Split(input, " ")
	if len(xInput) != 4 {
		return "", Node{}, fmt.Errorf("input '%s' did not split into 4 tokens", input)
	}

	thisAddress := xInput[0]

	leftAddress := strings.TrimPrefix(xInput[2], "(")
	leftAddress = strings.TrimSuffix(leftAddress, ",")

	node.left = Address(leftAddress)

	rightAddress := strings.TrimSuffix(xInput[3], ")")

	node.right = Address(rightAddress)

	return thisAddress, node, nil
}

func (p Puzzle) getToZZZ() (int, error) {
	if len(p.nodes) == 0 {
		return 0, nil
	}

	stepsTaken := 0

	var current Address = "AAA"

	currentNode, aaaNodeExists := p.nodes[current]
	if !aaaNodeExists {
		return 0, fmt.Errorf("the '%s' node does not exist", current)
	}

	for current != "ZZZ" {
		var nextNode Node
		var nodeExists bool

		switch direction := p.directions[stepsTaken%len(p.directions)]; direction {
		case Left:
			nextNode, nodeExists = p.nodes[currentNode.left]
			if !nodeExists {
				return 0, fmt.Errorf("next node '%s' to the left of '%#v' does not exist", currentNode.left, currentNode)
			}

			current = currentNode.left

		case Right:
			nextNode, nodeExists = p.nodes[currentNode.right]
			if !nodeExists {
				return 0, fmt.Errorf("next node '%s' to the right of '%#v' does not exist", currentNode.left, currentNode)
			}

			current = currentNode.right

		default:
			return 0, fmt.Errorf("direction '%#v' from current node '%#v' unknown", direction, currentNode)
		}

		stepsTaken++
		currentNode = nextNode
	}

	return stepsTaken, nil
}

func Part1(inputLines []string) (int, error) {
	puzzle, err := parseLines(inputLines)
	if err != nil {
		return 0, err
	}

	result, err := puzzle.getToZZZ()
	if err != nil {
		return 0, err
	}

	return result, nil
}
