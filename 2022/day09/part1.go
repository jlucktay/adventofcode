package day09

func TailVisitAtLeastOnce(input string) (int, error) {
	r := NewRope()

	if err := r.ParseCommands(input); err != nil {
		return 0, err
	}

	return r.TailVisitCount(), nil
}
