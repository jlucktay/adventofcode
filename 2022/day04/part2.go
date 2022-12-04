package day04

func SectionIDOverlap(input string) (int, error) {
	return sectionIDAnalysis(assignmentsOverlap, input)
}

func assignmentsOverlap(left, right []int) bool {
	if left[1] >= right[0] && right[1] >= left[0] {
		return true
	}

	return false
}
