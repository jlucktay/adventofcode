package day02

func CubeConundrumPartTwo(input string) (int, error) {
	games, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	runningTotal := 0

	for _, game := range games {
		runningTotal += game.FewestPossible()
	}

	return runningTotal, nil
}

func (gc GameCubes) FewestPossible() int {
	fewest := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, game := range gc {
		for colour, count := range game.colours {
			if fewest[colour] < count {
				fewest[colour] = count
			}
		}
	}

	return fewest["red"] * fewest["green"] * fewest["blue"]
}
