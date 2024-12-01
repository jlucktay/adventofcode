package crunchy

// AbsDiff returns the absolute difference between two ints.
func AbsDiff(x, y int) int {
	if x > y {
		return x - y
	}

	return y - x
}
