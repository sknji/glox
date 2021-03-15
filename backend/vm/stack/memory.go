package stack

const DefaultCapacity = 8

func GrowCapacity(cap int) int {
	if cap < DefaultCapacity {
		return DefaultCapacity
	}

	return cap * 2
}
