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
	self, left, right Address
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

		newNode, err := parseNode(inputLines[ilIndex])
		if err != nil {
			return Puzzle{}, err
		}

		puzzle.nodes[Address(newNode.self)] = newNode
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

func parseNode(input string) (Node, error) {
	node := Node{}

	xInput := strings.Split(input, " ")
	if len(xInput) != 4 {
		return Node{}, fmt.Errorf("input '%s' did not split into 4 tokens", input)
	}

	node.self = Address(xInput[0])

	leftAddress := strings.TrimPrefix(xInput[2], "(")
	leftAddress = strings.TrimSuffix(leftAddress, ",")

	node.left = Address(leftAddress)

	rightAddress := strings.TrimSuffix(xInput[3], ")")

	node.right = Address(rightAddress)

	return node, nil
}

func (p Puzzle) getToZZZ(ghostMode bool) (int, error) {
	if len(p.nodes) == 0 {
		return 0, nil
	}

	ghosts := make(Ghosts, 0)

	if ghostMode {
		for key := range p.nodes {
			if strings.HasSuffix(string(key), "A") {
				ghosts = append(ghosts, Ghost{current: key})
			}
		}

		if len(ghosts) == 0 {
			return 0, fmt.Errorf("no nodes ending with 'A' were found")
		}
	} else {
		ghosts = append(ghosts, Ghost{current: "AAA"})
	}

	for i := 0; i < len(ghosts); i++ {

		for !strings.HasSuffix(string(ghosts[i].current), "Z") {

			direction := p.directions[ghosts[i].stepsTaken%len(p.directions)]

			current, currentExists := p.nodes[ghosts[i].current]
			if !currentExists {
				return 0, fmt.Errorf("ghost #%d should be on the '%s' node but it does not exist", i, ghosts[i].current)
			}

			switch direction {
			case Left:
				if _, nodeExists := p.nodes[current.left]; !nodeExists {
					return 0, fmt.Errorf("next node '%s' to the left of '%#v' does not exist", current.left, current)
				}

				ghosts[i].current = current.left

			case Right:
				if _, nodeExists := p.nodes[current.right]; !nodeExists {
					return 0, fmt.Errorf("next node '%s' to the right of '%#v' does not exist", current.right, current)
				}

				ghosts[i].current = current.right

			default:
				return 0, fmt.Errorf("direction '%#v' from current node '%#v' unknown", direction, current)
			}

			ghosts[i].stepsTaken++
		}

		println("ghost", i, "took", ghosts[i].stepsTaken, "steps in total and finished at", ghosts[i].current)
	}

	stepsTaken := make([]int, 0)

	for i := 0; i < len(ghosts); i++ {
		stepsTaken = append(stepsTaken, ghosts[i].stepsTaken)
	}

	return lcm(stepsTaken...), nil
}

func Part1(inputLines []string) (int, error) {
	puzzle, err := parseLines(inputLines)
	if err != nil {
		return 0, err
	}

	result, err := puzzle.getToZZZ(false)
	if err != nil {
		return 0, err
	}

	return result, nil
}
