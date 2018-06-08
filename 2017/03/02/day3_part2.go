// Package main for Advent of Code 2017, day 3, part 2
// http://adventofcode.com/2017/day/3
package main

import (
	"fmt"

	"github.com/cavaliercoder/go-abs"
)

type direction uint8

const (
	right direction = iota
	up    direction = iota
	left  direction = iota
	down  direction = iota
)

type spiralNodeCoords struct {
	x, y int64
}

type spiralNode struct {
	parent *spiral
	coords spiralNodeCoords
	next   *spiralNode
	value  uint
}

/*
root:			(pointer to) the first node
last:			(pointer to) the last node
dir:			direction in which to add the next node
edgeDistance:	number of nodes already traversed on the current edge
sideLength:		the maximum number of nodes to add, before we need to change direction
firstChange:	side lengths come in pairs, so are we on the first of the two with this side length, or the second?
nodeMap:		a map of all existing nodes that have been added to this spiral
*/
type spiral struct {
	root, last               *spiralNode
	dir                      direction
	edgeDistance, sideLength uint
	firstChange              bool
	nodeMap                  map[spiralNodeCoords]*spiralNode
}

func (s *spiral) Init() {
	s.root = nil
	s.last = nil
	s.dir = right
	s.edgeDistance = 0
	s.sideLength = 1
	s.firstChange = true
	s.nodeMap = make(map[spiralNodeCoords]*spiralNode)
}

func (s *spiral) Add() {
	if s.root == nil {
		newNode := &spiralNode{s, spiralNodeCoords{0, 0}, nil, 1}
		s.root = newNode
		s.last = s.root
		s.nodeMap[newNode.coords] = newNode

		return
	}

	newNode := &spiralNode{s, s.last.nextCoords(), nil, 0}
	newNode.value = newNode.sumNeighbours()

	s.last.next = newNode
	s.last = s.last.next
	s.nodeMap[newNode.coords] = newNode

	if s.edgeDistance++; s.edgeDistance >= s.sideLength {
		s.turn()
	}
}

// Manhattan distance will be the sum of the absolutes of the last node's coordinates.
func (s *spiral) Manhattan() uint64 {
	return uint64(abs.WithTwosComplement(s.last.coords.x) + abs.WithTwosComplement(s.last.coords.y))
}

// Generates co-ordinates for all possible neighbours
func (snc spiralNodeCoords) neighbours() [8]spiralNodeCoords {
	return [...]spiralNodeCoords{
		{snc.x, snc.y + 1},
		{snc.x + 1, snc.y + 1},
		{snc.x + 1, snc.y},
		{snc.x + 1, snc.y - 1},
		{snc.x, snc.y - 1},
		{snc.x - 1, snc.y - 1},
		{snc.x - 1, snc.y},
		{snc.x - 1, snc.y + 1},
	}
}

// Calculates the sum of all actual neighbours' values
func (sn spiralNode) sumNeighbours() (result uint) {
	for _, neighbour := range sn.coords.neighbours() {
		if node, ok := sn.parent.nodeMap[neighbour]; ok {
			result += node.value
		}
	}

	return
}

func (sn spiralNode) String() string {
	return fmt.Sprintf("%s %v", sn.coords, sn.value)
}

func (snc spiralNodeCoords) String() string {
	return fmt.Sprintf("[%d,%d]", snc.x, snc.y)
}

func (s *spiral) turn() {
	s.edgeDistance = 0

	if s.firstChange {
		s.firstChange = false
	} else {
		s.firstChange = true
		s.sideLength++
	}

	switch s.dir {
	case right:
		s.dir = up
	case up:
		s.dir = left
	case left:
		s.dir = down
	default: //down
		s.dir = right
	}
}

func (sn spiralNode) nextCoords() spiralNodeCoords {
	switch sn.parent.dir {
	case right:
		return spiralNodeCoords{sn.coords.x + 1, sn.coords.y}
	case up:
		return spiralNodeCoords{sn.coords.x, sn.coords.y + 1}
	case left:
		return spiralNodeCoords{sn.coords.x - 1, sn.coords.y}
	default: //down
		return spiralNodeCoords{sn.coords.x, sn.coords.y - 1}
	}
}
