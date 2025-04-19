package day08

func HighestScenicScore(input string) (int, error) {
	ttg, err := NewTallTreeGrid(input)
	if err != nil {
		return 0, err
	}

	highestScore := 0

	for i := range len(ttg.trees) {
		for j := range len(ttg.trees[i]) {
			currentTreeScore := ttg.lookUpFrom(j, i) * ttg.lookRightFrom(j, i) * ttg.lookDownFrom(j, i) *
				ttg.lookLeftFrom(j, i)

			if currentTreeScore > highestScore {
				highestScore = currentTreeScore
			}
		}
	}

	return highestScore, nil
}

func (ttg *TallTreeGrid) lookUpFrom(x, y int) int {
	if y == 0 {
		return 0
	}

	heightOfStartingTree := ttg.trees[y][x]
	countVisibleTrees := 0

	for i := y - 1; i >= 0; i-- {
		countVisibleTrees++

		if heightOfStartingTree <= ttg.trees[i][x] {
			break
		}
	}

	return countVisibleTrees
}

func (ttg *TallTreeGrid) lookRightFrom(x, y int) int {
	if x == len(ttg.trees[y])-1 {
		return 0
	}

	heightOfStartingTree := ttg.trees[y][x]
	countVisibleTrees := 0

	for i := x + 1; i < len(ttg.trees[y]); i++ {
		countVisibleTrees++

		if heightOfStartingTree <= ttg.trees[y][i] {
			break
		}
	}

	return countVisibleTrees
}

func (ttg *TallTreeGrid) lookDownFrom(x, y int) int {
	if y == len(ttg.trees)-1 {
		return 0
	}

	heightOfStartingTree := ttg.trees[y][x]
	countVisibleTrees := 0

	for i := y + 1; i < len(ttg.trees); i++ {
		countVisibleTrees++

		if heightOfStartingTree <= ttg.trees[i][x] {
			break
		}
	}

	return countVisibleTrees
}

func (ttg *TallTreeGrid) lookLeftFrom(x, y int) int {
	if x == 0 {
		return 0
	}

	heightOfStartingTree := ttg.trees[y][x]
	countVisibleTrees := 0

	for i := x - 1; i >= 0; i-- {
		countVisibleTrees++

		if heightOfStartingTree <= ttg.trees[y][i] {
			break
		}
	}

	return countVisibleTrees
}
