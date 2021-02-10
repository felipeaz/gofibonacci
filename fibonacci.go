package fibonacci

// Sequence generates fibonacci with ${maxNums} elements.
func Sequence(maxNums uint) (sequence []int) {
	switch {
	case maxNums == 0:
		return []int{}
	case maxNums == 1:
		return []int{0}
	case maxNums == 2:
		return []int{0, 1}
	default:
		sequence = make([]int, int(maxNums))
		sequence[0], sequence[1] = 0, 1
	}

	for i := 1; i < int(maxNums)-1; i++ {
		sequence[i+1] = sequence[i] + sequence[i-1]
	}

	return
}

// NumberAtPosition returns the number at the given position on fibonacci sequence.
func NumberAtPosition(position uint) int {
	switch {
	case position == 0 || position == 1:
		return 0
	case position == 2:
		return 1
	}

	fibonacci := Sequence(position)

	return fibonacci[position-1]
}
