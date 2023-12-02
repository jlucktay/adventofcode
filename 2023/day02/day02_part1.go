package day02

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

func parseInput(input string) (map[int]GameCubes, error) {
	lines := strings.Split(input, "\n")
	result := make(map[int]GameCubes)

	for _, line := range lines {
		if line == "" {
			continue
		}

		games := strings.Split(line, ":")

		if len(games) != 2 {
			slog.Warn("expecting two parts after splitting on colon; continuing",
				slog.String("line", line),
				slog.String("games", fmt.Sprintf("%#v", games)))

			continue
		}

		var gameNumber int

		scanned, err := fmt.Sscanf(games[0], "Game %d", &gameNumber)
		if err != nil {
			return nil, err
		}

		slog.Debug("scanned numbers",
			slog.Int("count", scanned),
			slog.Int("game_number", gameNumber))

		if scanned != 1 {
			return nil, fmt.Errorf("expecting 1 scanned game number, got %d", scanned)
		}

		slog.Debug("remainder of line",
			slog.Int("len", len(games)),
			slog.String("key", games[1]))

		game := GameCubes{}

		cubeSets := strings.Split(games[1], ";")

		for _, cubeSet := range cubeSets {
			cubeSet = strings.TrimSpace(cubeSet)

			slog.Debug("cube set",
				slog.String("cube_set", cubeSet))

			gc := NewGameCube()

			game = append(game, gc)

			cubeColours := strings.Split(cubeSet, ",")

			for _, cubeColour := range cubeColours {
				cubeColour = strings.TrimSpace(cubeColour)

				slog.Debug("cube colour",
					slog.String("cube_colour", cubeColour))

				colourCount := strings.Split(cubeColour, " ")
				if len(colourCount) != 2 {
					slog.Warn("expecting two fields after split of cube colour;continuing",
						slog.String("cube_colour", cubeColour))

					continue
				}

				if err := gc.SetColour(colourCount[1], colourCount[0]); err != nil {
					slog.Error("setting colour",
						slog.Any("err", err))

					continue
				}
			}

			slog.Debug("game cube",
				slog.String("key", fmt.Sprintf("%+v", gc)))
		}

		result[gameNumber] = game
	}

	return result, nil
}

func CubeConundrum(input string) (int, error) {
	cubeColourLimitations := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	possibleGames := 0

	games, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	for gameNumber, game := range games {
		slog.Debug("game",
			slog.Int("game_number", gameNumber),
			slog.String("key", fmt.Sprintf("%+v", game)))

		if game.CheckPossible(cubeColourLimitations) {
			possibleGames += gameNumber
		}
	}

	return possibleGames, nil
}

type GameCube struct {
	colours map[string]int
}

func NewGameCube() GameCube {
	return GameCube{
		colours: make(map[string]int),
	}
}

func (gc GameCube) SetColour(colour, count string) error {
	parsedCount, err := strconv.Atoi(count)
	if err != nil {
		return err
	}

	gc.colours[colour] = parsedCount

	return nil
}

type GameCubes []GameCube

func (gc GameCubes) CheckPossible(limits map[string]int) bool {
	for _, cube := range gc {
		for colour, count := range cube.colours {
			if limits[colour] < count {
				slog.Warn("not possible",
					slog.String("cube", fmt.Sprintf("%+v", cube)),
					slog.String("limits", fmt.Sprintf("%+v", limits)),
				)

				return false
			}
		}
	}

	slog.Info("possible!",
		slog.String("limits", fmt.Sprintf("%+v", limits)),
	)

	return true
}
