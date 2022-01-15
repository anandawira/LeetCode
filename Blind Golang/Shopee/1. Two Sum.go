package shopee

func twoSum(nums []int, target int) []int {
	numIndex := make(map[int]int)

	for index, num := range nums {
		if v, ok := numIndex[target-num]; ok {
			return []int{index, v}
		}

		numIndex[num] = index
	}

	return nil
}
