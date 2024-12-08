/*
Scanning across the city, you find that there are actually many such antennas. Each antenna is tuned to a specific
frequency indicated by a single lowercase letter, uppercase letter, or digit. You create a map (your puzzle input) of
these antennas. For example:

............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............

The signal only applies its nefarious effect at specific antinodes based on the resonant frequencies of the antennas.
In particular, an antinode occurs at any point that is perfectly in line with two antennas of the same frequency - but
only when one of the antennas is twice as far away as the other. This means that for any pair of antennas with the same
frequency, there are two antinodes, one on either side of them.

So, for these two antennas with frequency a, they create the two antinodes marked with #:

..........
...#......
..........
....a.....
..........
.....a....
..........
......#...
..........
..........

Adding a third antenna with the same frequency creates several more antinodes. It would ideally add four antinodes, but
two are off the right side of the map, so instead it adds only two:

..........
...#......
#.........
....a.....
........a.
.....a....
..#.......
......#...
..........
..........

Antennas with different frequencies don't create antinodes; A and a count as different frequencies. However, antinodes
can occur at locations that contain antennas. In this diagram, the lone antenna with frequency capital A creates no
antinodes but has a lowercase-a-frequency antinode at its location:

..........
...#......
#.........
....a.....
........a.
.....a....
..#.......
......A...
..........
..........

The first example has antennas with two different frequencies, so the antinodes they create look like this, plus an
antinode overlapping the topmost A-frequency antenna:

......#....#
...#....0...
....#0....#.
..#....0....
....0....#..
.#....A.....
...#........
#......#....
........A...
.........A..
..........#.
..........#.

Because the topmost A-frequency antenna overlaps with a 0-frequency antinode, there are 14 total unique locations that
contain an antinode within the bounds of the map.

Calculate the impact of the signal. How many unique locations within the bounds of the map contain an antinode?
*/

package day08_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2024/day08"
)

func TestPart1(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   string
		want int
	}{
		"empty": {"", 0},
		"for example": {`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`, 14},
		"for these two antennas with frequency a, they create the two antinodes": {`..........
..........
..........
....a.....
..........
.....a....
..........
..........
..........
..........
`, 2},
		"Adding a third antenna with the same frequency creates several more antinodes": {`..........
..........
..........
....a.....
........a.
.....a....
..........
..........
..........
..........
`, 4},
	}

	for desc, testCase := range testCases {
		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := day08.Part1(testCase.in)
			is.NoErr(err)
			is.Equal(got, testCase.want)
		})
	}
}
