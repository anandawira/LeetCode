package array

func maxProduct(nums []int) int {
	minval := nums[0]
	maxval := nums[0]
	maxproduct := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			minval, maxval = maxval, minval
		}

		minval = min(nums[i], minval*nums[i])
		maxval = max(nums[i], maxval*nums[i])
		if maxval > maxproduct {
			maxproduct = maxval
		}
	}
	return maxproduct
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
