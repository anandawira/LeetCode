func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sums := make(map[int]int)
	var answer int

	for _, a := range nums1 {
		for _, b := range nums2 {
			sums[a+b]++
		}
	}

	for _, c := range nums3 {
		for _, d := range nums4 {
			answer += sums[0-c-d]
		}
	}
	return answer
}