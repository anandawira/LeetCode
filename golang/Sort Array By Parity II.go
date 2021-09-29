package sort_array_by_parity_ii

func sortArrayByParityII(nums []int) []int {
	// To store misplaced odds and evens number
	var (
		evenIndex = 0
		oddIndex  = 1
	)

	var result = make([]int, len(nums))

	for _, num := range nums {
		if num%2 == 0 {
			result[evenIndex] = num
			evenIndex += 2
		} else {
			result[oddIndex] = num
			oddIndex += 2
		}
	}

	return result
}

// func sortArrayByParityII(nums []int) []int {
// 	// To store misplaced odds and evens number
// 	var odds []int
// 	var evens []int

// 	for i, num := range nums {
// 		if isEven(i) != isEven(num) {
// 			if isEven(i) {
// 				odds = append(odds, i)
// 			} else {
// 				evens = append(evens, i)
// 			}
// 		}
// 	}

// 	// Swap misplaced odds and evens number pair
// 	for i := 0; i < len(odds); i++ {
// 		nums[odds[i]], nums[evens[i]] = nums[evens[i]], nums[odds[i]]
// 	}

// 	return nums
// }

// func sortArrayByParityII(nums []int) []int {
// 	var odds []int
// 	var evens []int

// 	for _, num := range nums {
// 		if num%2 == 0 {
// 			evens = append(evens, num)
// 		} else {
// 			odds = append(odds, num)
// 		}
// 	}

// 	var result []int

// 	for i, even := range evens {
// 		result = append(result, even, odds[i])
// 	}

// 	return result
// }
