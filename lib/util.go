package lib

func GrowCapacity(cap int) int {
	if cap < 8 {
		return 8
	}

	return cap * 2
}
