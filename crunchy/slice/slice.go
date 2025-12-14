// Package slice holds helpers for working with slices.
//
// The singular was used so as not to collide with the pluralised standard library package.
package slice

import (
	"slices"
)

// Union creates a new slice containing the union of the two input slices a and b.
// The ordering of the returned slice is randomised.
func Union[T comparable](a, b []T) []T {
	if a == nil && b == nil {
		return nil
	}

	if a == nil && b != nil {
		cpy := make([]T, len(b))
		copy(cpy, b)

		return cpy
	}

	if a != nil && b == nil {
		cpy := make([]T, len(a))
		copy(cpy, a)

		return cpy
	}

	set := make(map[T]struct{})

	for _, aa := range a {
		set[aa] = struct{}{}
	}

	for _, bb := range b {
		set[bb] = struct{}{}
	}

	result := make([]T, len(set))

	i := 0

	for k := range set {
		result[i] = k
		i++
	}

	return result
}

// Delete elements at all given indices, and return the modified slice.
// The indices can be given in any order, and any that are out of bounds will be ignored.
func Delete[T any](s []T, indices ...int) []T {
	slices.Sort(indices)
	indices = slices.Compact(indices)
	slices.Reverse(indices)

	for _, index := range indices {
		if index >= len(s) || index < 0 {
			continue
		}

		s = append(s[:index], s[index+1:]...)
	}

	return s
}
