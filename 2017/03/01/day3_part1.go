/*
http://adventofcode.com/2017/day/3
*/

package main

import (
	"fmt"
	"math"
)

type direction uint8

const (
	right direction = iota
	up    direction = iota
	left  direction = iota
	down  direction = iota
)

type spiralNodeCoords struct {
	x, y int
}

type spiralNode struct {
	coords spiralNodeCoords
	next   *spiralNode
	value  uint
}

/*
root:			(pointer to) the first node
last:			(pointer to) the last node
dir:			direction in which to add the next node
firstChange:	side lengths come in pairs, so are we on the first of the two with this side length, or the second?
edgeDistance:	number of nodes already traversed on the current edge
sideLength:		the maximum number of nodes to add, before we need to change direction
*/
type spiral struct {
	root, last               *spiralNode
	dir                      direction
	firstChange              bool
	edgeDistance, sideLength uint
}

func (s *spiral) Init() {
	s.root = nil
	s.last = nil
	s.dir = right
	s.firstChange = true
	s.edgeDistance = 0
	s.sideLength = 1
}

func (s *spiral) Add() {
	if s.root == nil {
		s.root = &spiralNode{
			spiralNodeCoords{0, 0},
			nil,
			1,
		}
		s.last = s.root

		return
	}

	s.last.next = &spiralNode{
		s.last.coords.next(s.dir),
		nil,
		s.last.value + 1,
	}

	s.last = s.last.next
	s.edgeDistance++

	if s.edgeDistance >= s.sideLength {
		s.edgeDistance = 0
		s.dir = s.dir.turn()

		if s.firstChange {
			s.firstChange = false
		} else {
			s.firstChange = true
			s.sideLength++
		}
	}

	// if s.last.coords.x > 0 && s.last.coords.x == s.last.coords.y*-1 {
	// 	fmt.Printf("%d..", s.Size())
	// }
}

func (s *spiral) Size() uint {
	if s.root == nil {
		return 0
	}

	return s.last.value
}

/*
Manhattan distance will be the sum of the absolutes of the last node's coordinates.
*/
func (s *spiral) Manhattan() uint {
	return uint(math.Abs(float64(s.last.coords.x)) + math.Abs(float64(s.last.coords.y)))
}

func (d direction) turn() direction {
	switch d {
	case right:
		return up
	case up:
		return left
	case left:
		return down
	default: //down
		return right
	}
}

func (snc spiralNodeCoords) next(dir direction) spiralNodeCoords {
	switch dir {
	case right:
		return spiralNodeCoords{snc.x + 1, snc.y}
	case up:
		return spiralNodeCoords{snc.x, snc.y + 1}
	case left:
		return spiralNodeCoords{snc.x - 1, snc.y}
	default: //down
		return spiralNodeCoords{snc.x, snc.y - 1}
	}
}

func main() {
	var spir spiral
	spir.Init()

	target := 277678

	for index := 0; index < target; index++ {
		spir.Add()
	}

	// fmt.Println("\nlast:", spir.last)
	fmt.Println("Manhattan distance for", target, "is", spir.Manhattan())
}
