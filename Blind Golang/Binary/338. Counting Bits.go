package binary

func countBits(n int) []int {
	result := make([]int, n+1)
	offset := 1

	for i := 1; i < n+1; i++ {
		if offset*2 == i {
			offset = i
		}

		result[i] = 1 + result[i-offset]
	}

	return result
}
