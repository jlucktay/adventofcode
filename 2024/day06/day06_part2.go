// Package day06 for Advent of Code 2024, day 6, part 2.
// https://adventofcode.com/2024/day/6
package day06

import (
	"errors"
	"fmt"
	"log/slog"
)

func Part2(input string) (int, error) {
	myMap, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	loopsDetected := 0

	var mapHeight, mapWidth int

	mapHeight = len(myMap.floorPlan)

	if mapHeight > 0 {
		mapWidth = len(myMap.floorPlan[0])
	}

	// Cycle across the floorplan, put a new obstruction on each 'Empty' spot, then run the loop detector.
	for rowNum := range mapHeight {
		for colNum := range mapWidth {
			slog.Debug("next candidate in cycle",
				slog.Int("x", colNum), slog.Int("y", rowNum))

			if myMap.floorPlan[rowNum][colNum] != Empty {
				slog.Debug("position is not empty, continuing",
					slog.Int("x", colNum), slog.Int("y", rowNum))

				continue
			}

			var errProtocol error

			myMap.floorPlan[rowNum][colNum] = Obstruction

			for errProtocol == nil {
				errProtocol = myMap.FollowProtocolDetectLoops()
			}

			if errors.Is(errProtocol, ErrLoopDetected) {
				slog.Debug("finished candidate's 'for' loop",
					slog.Any("errProtocol", errProtocol),
					slog.Int("x", colNum), slog.Int("y", rowNum))

				loopsDetected++
			}

			myMap, err = parseInput(input)
			if err != nil {
				return 0, err
			}
		}
	}

	return loopsDetected, nil
}

var ErrLoopDetected = errors.New("loop detected")

func (m *Map) FollowProtocolDetectLoops() error {
	slog.Debug("findGuard", slog.Int("x", m.guardX), slog.Int("y", m.guardY))

	if m.visited[[2]int{m.guardX, m.guardY}] == nil {
		m.visited[[2]int{m.guardX, m.guardY}] = make(map[MapPoint]struct{})
	}

	switch m.lookInFront() {
	case OutOfBounds:
		slog.Debug("case OutOfBounds")

		m.visited[[2]int{m.guardX, m.guardY}][Visited] = struct{}{}

		return ErrOutOfBounds

	case Obstruction:
		slog.Debug("case Obstruction")

		currentGuardFacing := m.floorPlan[m.guardY][m.guardX]

		m.floorPlan[m.guardY][m.guardX] = TurnRight(m.floorPlan[m.guardY][m.guardX])

		slog.Debug("turning because obstruction", slog.Int("x", m.guardX), slog.Int("y", m.guardY))

		m.visited[[2]int{m.guardX, m.guardY}][currentGuardFacing] = struct{}{}

	case Empty:
		slog.Debug("case Empty")

		// Current guard position becomes empty, and we save which way the guard was facing.
		currentGuardFacing := m.floorPlan[m.guardY][m.guardX]
		nextX, nextY := m.getInFront()

		if alreadyVisitedDirection, alreadyVisited := m.visited[[2]int{nextX, nextY}]; alreadyVisited {
			if _, overlap := alreadyVisitedDirection[currentGuardFacing]; overlap {
				return fmt.Errorf("%w: %d", ErrLoopDetected, len(m.visited))
			}
		}

		m.floorPlan[m.guardY][m.guardX] = Empty

		m.visited[[2]int{m.guardX, m.guardY}][currentGuardFacing] = struct{}{}

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
