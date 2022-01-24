package random

func findMaxConsecutiveOnes(nums []int) int {
	var max, count int

	for _, num := range nums {
		if num == 1 {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
	}
	return max
}
