package daily

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1

	var max int
	for l < r {
		amount := (r - l) * min(height[r], height[l])
		if amount > max {
			max = amount
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
