package day09

type Position struct {
	x, y int
}

func NewPosition(x, y int) Position {
	return Position{x, y}
}

func (p *Position) DistanceFrom(q Position) int {
	xDelta := abs(p.x - q.x)
	yDelta := abs(p.y - q.y)

	biggest := xDelta
	if yDelta > xDelta {
		biggest = yDelta
	}

	return biggest
}

type Direction string

const (
	DirUp    Direction = "up"
	DirRight Direction = "right"
	DirDown  Direction = "down"
	DirLeft  Direction = "left"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
