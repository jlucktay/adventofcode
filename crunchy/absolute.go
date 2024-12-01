package crunchy

func AbsoluteDiff(x, y int) int {
	if x > y {
		return x - y
	}

	return y - x
}
