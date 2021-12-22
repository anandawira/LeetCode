package array

func twoSum(nums []int, target int) []int {
	checkedNumbers := make(map[int]int)

	for index, cur := range nums {
		diff := target - cur

		if v, ok := checkedNumbers[diff]; ok {
			return []int{index, v}
		}

		checkedNumbers[cur] = index
	}

	return nil
}
