package random

func twoSum(nums []int, target int) []int {
	previousNums := make(map[int]int)

	for firstIndex, num := range nums {
		diff := target - num

		secondIndex, ok := previousNums[diff]

		if ok {
			return []int{firstIndex, secondIndex}
		}

		previousNums[num] = firstIndex
	}

	return nil
}
