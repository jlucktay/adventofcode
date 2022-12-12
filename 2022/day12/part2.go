package day12

func Part2(input string) (int, error) {
	length, err := bfs(ParseWorld(input), true)
	if err != nil {
		return 0, err
	}

	return length, nil
}
