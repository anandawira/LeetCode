package array

func maxSubArray(nums []int) int {
	maxSums := nums
	for i := 1; i < len(nums); i++ {
		if maxSums[i-1] > 0 {
			maxSums[i] += maxSums[i-1]
		}
	}
	return maxOfSlice(maxSums)
}

func maxOfSlice(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if cur := nums[i]; cur > max {
			max = cur
		}
	}

	return max
}
