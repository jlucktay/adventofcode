package day11

func TenThousandRoundsOfMonkeyBusiness(input string) (int, error) {
	return SimulateRoundsOfMonkeyBusiness(input, 10000, true)
}

// Big thanks to Reddit, Stack Overflow, and Siong-Ui Te on GitHub here:
// https://www.reddit.com/r/adventofcode/comments/zifqmh/2022_day_11_solutions/?sort=old
// https://stackoverflow.com/questions/147515/least-common-multiple-for-3-or-more-numbers
// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// Greatest common divisor (GCD) via Euclidean algorithm.
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

// Least common multiple (LCM) via GCD.
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := range len(integers) {
		result = lcm(result, integers[i])
	}

	return result
}
