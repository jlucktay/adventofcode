package day02

import (
	"testing"

	"github.com/matryer/is"
)

func TestReportAscOrDesc(t *testing.T) {
	is := is.New(t)

	is.True(reportAscOrDesc([]int{1, 2, 3}))
	is.True(reportAscOrDesc([]int{3, 2, 1}))
	is.True(!reportAscOrDesc([]int{3, 2, 3}))
}

func TestReportStrictlyAsc(t *testing.T) {
	is := is.New(t)

	is.True(reportStrictlyAsc([]int{1, 2, 3}))
	is.True(!reportStrictlyAsc([]int{3, 2, 1}))
	is.True(!reportStrictlyAsc([]int{3, 2, 3}))
	is.True(!reportStrictlyAsc([]int{2, 2, 3}))
}

func TestReportStrictlyDesc(t *testing.T) {
	is := is.New(t)

	is.True(!reportStrictlyDesc([]int{1, 2, 3}))
	is.True(reportStrictlyDesc([]int{3, 2, 1}))
	is.True(!reportStrictlyDesc([]int{3, 2, 3}))
	is.True(!reportStrictlyDesc([]int{2, 2, 3}))
}

func TestReportHasDuplicates(t *testing.T) {
	is := is.New(t)

	is.True(!reportHasDuplicates([]int{1, 2, 3}))
	is.True(reportHasDuplicates([]int{1, 1, 2, 2, 3, 3}))
}

func TestReportGapsBetween1And3(t *testing.T) {
	is := is.New(t)

	is.True(reportGapsBetween1And3([]int{1, 2, 3}))
	is.True(!reportGapsBetween1And3([]int{1, 5, 3}))

	is.True(!reportGapsBetween1And3([]int{1, 2, 7, 8, 9}))
	is.True(!reportGapsBetween1And3([]int{9, 7, 6, 2, 1}))
	is.True(!reportGapsBetween1And3([]int{8, 6, 4, 4, 1}))
}

func TestReportRemoveFirstUnsafe(t *testing.T) {
	is := is.New(t)

	is.Equal(reportRemoveFirstUnsafe([]int{1, 1, 2, 4, 5}), []int{1, 2, 4, 5})
	is.Equal(reportRemoveFirstUnsafe([]int{1, 2, 2, 4, 4}), nil)
	is.Equal(reportRemoveFirstUnsafe([]int{1, 3, 2, 4, 5}), []int{1, 2, 4, 5})
	is.Equal(reportRemoveFirstUnsafe([]int{8, 6, 4, 4, 1}), []int{8, 6, 4, 1})
}
