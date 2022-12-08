package day08

func (ttg *TallTreeGrid) VisibleFromOutsideGrid() int {
	visibleCount := 0

	// Every tree around the outside is automatically visible.
	horizontalGridSize := len(ttg.trees[0])
	verticalGridSize := len(ttg.trees)
	perimeter := horizontalGridSize*2 + (verticalGridSize-2)*2
	visibleCount += perimeter

	for i := 1; i < len(ttg.trees)-1; i++ {
		for j := 1; j < len(ttg.trees[i])-1; j++ {
			// Look all four ways
			if ttg.lookDownAt(j, i) || ttg.lookLeftAt(j, i) || ttg.lookUpAt(j, i) || ttg.lookRightAt(j, i) {
				visibleCount++
			}
		}
	}

	return visibleCount
}

func (ttg *TallTreeGrid) lookDownAt(x, y int) bool {
	heightOfTargetTree := ttg.trees[y][x]

	for i := 0; i < y; i++ {
		if ttg.trees[i][x] >= heightOfTargetTree {
			return false
		}
	}

	return true
}

func (ttg *TallTreeGrid) lookLeftAt(x, y int) bool {
	heightOfTargetTree := ttg.trees[y][x]

	for i := len(ttg.trees[y]) - 1; i > x; i-- {
		if ttg.trees[y][i] >= heightOfTargetTree {
			return false
		}
	}

	return true
}

func (ttg *TallTreeGrid) lookUpAt(x, y int) bool {
	heightOfTargetTree := ttg.trees[y][x]

	for i := len(ttg.trees) - 1; i > y; i-- {
		if ttg.trees[i][x] >= heightOfTargetTree {
			return false
		}
	}

	return true
}

func (ttg *TallTreeGrid) lookRightAt(x, y int) bool {
	heightOfTargetTree := ttg.trees[y][x]

	for i := 0; i < x; i++ {
		if ttg.trees[y][i] >= heightOfTargetTree {
			return false
		}
	}

	return true
}

func TreesVisibleFromOutsideGrid(input string) (int, error) {
	ttg, err := NewTallTreeGrid(input)

	return ttg.VisibleFromOutsideGrid(), err
}
