package day04

func SectionIDFullyContain(input string) (int, error) {
	return sectionIDAnalysis(assignmentsFullyContain, input)
}

func assignmentsFullyContain(left, right []int) bool {
	if left[0] <= right[0] && left[1] >= right[1] {
		return true
	}

	if right[0] <= left[0] && right[1] >= left[1] {
		return true
	}

	return false
}
