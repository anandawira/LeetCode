package daily

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[l] {
			if target > nums[mid] || target < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if target < nums[mid] || target > nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
