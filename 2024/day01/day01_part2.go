package day01

func ListSimilarity(input string) (int, error) {
	lists, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := 0

	rhsCounts := map[int]int{}

	for index := range lists.right {
		rhsCounts[lists.right[index]]++
	}

	for index := range lists.left {
		left := lists.left[index]

		if right, ok := rhsCounts[left]; ok {
			result += left * right
		}
	}

	return result, nil
}
