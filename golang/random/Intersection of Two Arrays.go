package random

func intersection(nums1 []int, nums2 []int) []int {
	seen := make(map[int]bool)
	for _, num := range nums1 {
		seen[num] = true
	}

	var result []int

	for _, num := range nums2 {
		if exist := seen[num]; exist {
			seen[num] = false
			result = append(result, num)
		}
	}

	return result
}
