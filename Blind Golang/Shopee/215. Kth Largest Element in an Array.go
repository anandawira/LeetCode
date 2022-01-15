package shopee

func findKthLargest(nums []int, k int) int {
	start := 0
	end := len(nums) - 1
	targetPos := len(nums) - k

	for {
		pivotIndex := partition(nums, start, end)
		if pivotIndex == targetPos {
			return nums[pivotIndex]
		} else if pivotIndex < targetPos {
			start = pivotIndex + 1
		} else {
			end = pivotIndex - 1
		}
	}

}

func partition(nums []int, start, end int) int {
	pivot := nums[start]

	l, r := start+1, end

	for l <= r {
		for l <= r && nums[l] <= pivot {
			l++
		}
		for l <= r && nums[r] > pivot {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}

	nums[r], nums[start] = nums[start], nums[r]

	return r
}
