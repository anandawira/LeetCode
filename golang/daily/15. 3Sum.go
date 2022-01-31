package daily

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		target := nums[i] * -1

		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := nums[l] + nums[r]
			if sum > target {
				r--
			} else if sum < target {
				l++
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				r--
				l++

				for nums[l] == nums[l-1] && l < r {
					l++
				}

				for nums[r] == nums[r+1] && l < r {
					r--
				}
			}
		}
	}
	return result
}
