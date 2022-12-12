package day12

func Part1(input string) (int, error) {
	length, err := bfs(ParseWorld(input), false)
	if err != nil {
		return 0, err
	}

	return length, nil
}
