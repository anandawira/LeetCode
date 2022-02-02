package dp

func lengthOfLIS(nums []int) int {
	maxLength := make([]int, len(nums))
	maxLength[0] = 1
	result := 1
	for i := 1; i < len(nums); i++ {
		maxLength[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && maxLength[j]+1 > maxLength[i] {
				maxLength[i] = maxLength[j] + 1

				if maxLength[i] > result {
					result = maxLength[i]
				}
			}
		}
	}

	return result
}
