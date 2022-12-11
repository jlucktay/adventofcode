package day11_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day11"
)

func TestTenThousandRoundsOfMonkeyBusiness(t *testing.T) {
	is := is.New(t)

	actual, err := day11.TenThousandRoundsOfMonkeyBusiness(INPUT)
	is.NoErr(err)
	is.Equal(2_713_310_158, actual)
}
