package day09

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Rope struct {
	head, tail Position
	tailVisits map[Position]struct{}
}

func NewRope() Rope {
	return Rope{
		head:       NewPosition(0, 0),
		tail:       NewPosition(0, 0),
		tailVisits: make(map[Position]struct{}),
	}
}

func (r *Rope) TailVisitCount() int {
	return len(r.tailVisits)
}

func (r *Rope) MoveHead(dir Direction) {
	switch dir {
	case DirUp:
		r.head.y++
	case DirRight:
		r.head.x++
	case DirDown:
		r.head.y--
	case DirLeft:
		r.head.x--
	}
}

func (r *Rope) TailMustMove() bool {
	dist := r.tail.DistanceFrom(r.head)
	return dist > 1
}

func (r *Rope) TailCatchUp() {
	r.tailVisits[Position{r.tail.x, r.tail.y}] = struct{}{}

	defer func() {
		r.tailVisits[Position{r.tail.x, r.tail.y}] = struct{}{}
	}()

	if r.head.x == r.tail.x {
		if r.head.y > r.tail.y {
			r.tail.y++
		} else {
			r.tail.y--
		}

		return
	}

	if r.head.y == r.tail.y {
		if r.head.x > r.tail.x {
			r.tail.x++
		} else {
			r.tail.x--
		}

		return
	}

	if r.head.x > r.tail.x && r.head.y > r.tail.y {
		r.tail.x++
		r.tail.y++
	} else if r.head.x > r.tail.x && r.head.y < r.tail.y {
		r.tail.x++
		r.tail.y--
	} else if r.head.x < r.tail.x && r.head.y < r.tail.y {
		r.tail.x--
		r.tail.y--
	} else if r.head.x < r.tail.x && r.head.y > r.tail.y {
		r.tail.x--
		r.tail.y++
	}
}

func (r *Rope) ParseCommands(input string) error {
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
			return err
		}

		for i := 0; i < int(distance); i++ {
			r.MoveHead(dir)

			if r.TailMustMove() {
				r.TailCatchUp()
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanning input: %v", err)
	}

	return nil
}

func TailVisitAtLeastOnce(input string) (int, error) {
	r := NewRope()

	if err := r.ParseCommands(input); err != nil {
		return 0, err
	}

	return r.TailVisitCount(), nil
}
