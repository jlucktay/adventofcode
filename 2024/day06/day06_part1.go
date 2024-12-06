// Package day06 for Advent of Code 2024, day 6, part 1.
// https://adventofcode.com/2024/day/6
package day06

import (
	"errors"
	"log/slog"
	"strings"
	"unicode"

	"github.com/orsinium-labs/enum"
)

func Part1(input string) (int, error) {
	myMap, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	var errProtocol error

	for errProtocol == nil {
		errProtocol = myMap.FollowProtocol()
	}

	return len(myMap.visited), nil
}

type MapPoint enum.Member[rune]

var (
	Empty       = MapPoint{'.'}
	Obstruction = MapPoint{'#'}

	GuardNorth = MapPoint{'^'}
	GuardEast  = MapPoint{'>'}
	GuardSouth = MapPoint{'v'}
	GuardWest  = MapPoint{'<'}

	OutOfBounds = MapPoint{unicode.MaxRune}

	MapPoints      = enum.New(Empty, Obstruction, GuardNorth, GuardEast, GuardSouth, GuardWest, OutOfBounds)
	GuardMapPoints = enum.New(GuardNorth, GuardEast, GuardSouth, GuardWest)
)

func TurnRight(in MapPoint) MapPoint {
	switch in {
	case GuardNorth:
		return GuardEast
	case GuardEast:
		return GuardSouth
	case GuardSouth:
		return GuardWest
	case GuardWest:
		return GuardNorth
	default:
		return in
	}
}

type Map struct {
	guardX, guardY int

	floorPlan [][]MapPoint

	visited map[[2]int]struct{}
}

func (m Map) String() string {
	output := strings.Builder{}

	for _, mapRow := range m.floorPlan {
		for _, mapPoint := range mapRow {
			_, _ = output.WriteRune(mapPoint.Value)
		}

		output.WriteString("\n")
	}

	return output.String()
}

var (
	ErrOutOfBounds = errors.New("out of bounds")
	ErrProtocol    = errors.New("could not follow protocol")
)

func (m *Map) FollowProtocol() error {
	slog.Debug("findGuard", slog.Int("x", m.guardX), slog.Int("y", m.guardY))

	switch m.lookInFront() {
	case OutOfBounds:
		slog.Debug("case OutOfBounds")

		m.visited[[2]int{m.guardX, m.guardY}] = struct{}{}

		return ErrOutOfBounds

	case Obstruction:
		slog.Debug("case Obstruction")

		m.floorPlan[m.guardY][m.guardX] = TurnRight(m.floorPlan[m.guardY][m.guardX])

	case Empty:
		slog.Debug("case Empty")

		// Current guard position becomes empty, and we save which way the guard was facing.
		currentGuardFacing := m.floorPlan[m.guardY][m.guardX]
		nextX, nextY := m.getInFront()
		m.floorPlan[m.guardY][m.guardX] = Empty
		m.visited[[2]int{m.guardX, m.guardY}] = struct{}{}

		// Next guard position becomes the guard rune facing the same direction.
		m.floorPlan[nextY][nextX] = currentGuardFacing
		m.guardX = nextX
		m.guardY = nextY

	default:
		slog.Debug("default")

		return ErrProtocol
	}

	return nil
}

func (m *Map) getInFront() (int, int) {
	if m.guardY < 0 {
		return -1, -1
	}

	if m.guardY > len(m.floorPlan) {
		return -1, -1
	}

	if m.guardX < 0 {
		return -1, -1
	}

	if len(m.floorPlan) > 0 && m.guardX > len(m.floorPlan[0]) {
		return -1, -1
	}

	currentFacing := m.floorPlan[m.guardY][m.guardX]

	var nextX, nextY int

	switch currentFacing {
	case GuardNorth:
		nextX = m.guardX
		nextY = m.guardY - 1
	case GuardEast:
		nextX = m.guardX + 1
		nextY = m.guardY
	case GuardSouth:
		nextX = m.guardX
		nextY = m.guardY + 1
	case GuardWest:
		nextX = m.guardX - 1
		nextY = m.guardY
	}

	return nextX, nextY
}

func (m *Map) lookInFront() MapPoint {
	nextX, nextY := m.getInFront()

	if nextX < 0 || nextY < 0 || nextX >= len(m.floorPlan[0]) || nextY >= len(m.floorPlan) {
		return OutOfBounds
	}

	return m.floorPlan[nextY][nextX]
}
