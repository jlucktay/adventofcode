package crunchy

import "golang.org/x/exp/constraints"

// LCM calculates the [least common multiple] of all given integers.
//
// [least common multiple]: https://en.wikipedia.org/wiki/Least_common_multiple
func LCM[T constraints.Integer](x ...T) T {
	if len(x) == 0 {
		return 0
	} else if len(x) == 1 {
		return x[0]
	} else if len(x) == 2 {
		return x[0] * x[1] / GCD(x[0], x[1])
	}

	return LCM(x[0], LCM(x[1:]...))
}
