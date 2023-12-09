package crunchy

import "golang.org/x/exp/constraints"

// GCD calculates the [greatest common divisor] of the two given integers.
//
// [greatest common divisor]: https://en.wikipedia.org/wiki/Greatest_common_divisor
func GCD[T constraints.Integer](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
