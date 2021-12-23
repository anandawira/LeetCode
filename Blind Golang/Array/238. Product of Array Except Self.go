package array

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	prod := 1
	for i, n := range nums {
		result[i] = prod
		prod *= n
	}
	prod = 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= prod
		prod *= nums[i]
	}
	return result
}
