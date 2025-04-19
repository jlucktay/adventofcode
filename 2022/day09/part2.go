package day09

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type LongRope struct {
	Rope
	knots [8]Position
}

func NewLongRope() LongRope {
	return LongRope{
		Rope: NewRope(),
		knots: [8]Position{
			{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0},
		},
	}
}

func (lr *LongRope) KnotMustMove(index int) bool {
	thisKnot, knotBeingFollowed, err := lr.getKnotAndFollowed(index)
	if err != nil {
		return false
	}

	return thisKnot.DistanceFrom(*knotBeingFollowed) > 1
}

func (lr *LongRope) KnotCatchUp(index int) error {
	thisKnot, knotBeingFollowed, err := lr.getKnotAndFollowed(index)
	if err != nil {
		return err
	}

	if *thisKnot == lr.tail {
		lr.tailVisits[Position{lr.tail.x, lr.tail.y}] = struct{}{}

		defer func() {
			lr.tailVisits[Position{lr.tail.x, lr.tail.y}] = struct{}{}
		}()
	}

	if knotBeingFollowed.x == thisKnot.x {
		if knotBeingFollowed.y > thisKnot.y {
			thisKnot.y++
		} else {
			thisKnot.y--
		}

		return nil
	}

	if knotBeingFollowed.y == thisKnot.y {
		if knotBeingFollowed.x > thisKnot.x {
			thisKnot.x++
		} else {
			thisKnot.x--
		}

		return nil
	}

	if knotBeingFollowed.x > thisKnot.x && knotBeingFollowed.y > thisKnot.y {
		thisKnot.x++
		thisKnot.y++
	} else if knotBeingFollowed.x > thisKnot.x && knotBeingFollowed.y < thisKnot.y {
		thisKnot.x++
		thisKnot.y--
	} else if knotBeingFollowed.x < thisKnot.x && knotBeingFollowed.y < thisKnot.y {
		thisKnot.x--
		thisKnot.y--
	} else if knotBeingFollowed.x < thisKnot.x && knotBeingFollowed.y > thisKnot.y {
		thisKnot.x--
		thisKnot.y++
	}

	return nil
}

func (lr *LongRope) getKnotAndFollowed(index int) (*Position, *Position, error) {
	var thisKnot, knotBeingFollowed *Position

	if index < 0 || index > 8 {
		return nil, nil, fmt.Errorf("index %d out of bounds", index)
	}

	switch index {
	case 0:
		thisKnot = &lr.knots[0]
		knotBeingFollowed = &lr.head
	case 8:
		thisKnot = &lr.tail
		knotBeingFollowed = &lr.knots[7]
	default:
		thisKnot = &lr.knots[index]
		knotBeingFollowed = &lr.knots[index-1]
	}

	return thisKnot, knotBeingFollowed, nil
}

func (lr *LongRope) ParseCommands(input string) error {
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		xLine := strings.Split(line, " ")

		if len(xLine) != 2 {
			return fmt.Errorf("bad input line: '%s'", line)
		}

		var dir Direction

		switch xLine[0] {
		case "U":
			dir = DirUp
		case "R":
			dir = DirRight
		case "D":
			dir = DirDown
		case "L":
			dir = DirLeft
		}

		distance, err := strconv.ParseInt(xLine[1], 10, 32)
		if err != nil {
			return fmt.Errorf("parsing integer: %w", err)
		}

		for range distance {
			lr.MoveHead(dir)

			for index := range lr.knots {
				if lr.KnotMustMove(index) {
					if err := lr.KnotCatchUp(index); err != nil {
						return err
					}
				}
			}

			if lr.KnotMustMove(len(lr.knots)) {
				if err := lr.KnotCatchUp(len(lr.knots)); err != nil {
					return err
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanning input: %w", err)
	}

	return nil
}

func LongTailVisitAtLeastOnce(input string) (int, error) {
	lr := NewLongRope()

	if err := lr.ParseCommands(input); err != nil {
		return 0, err
	}

	return lr.TailVisitCount(), nil
}
