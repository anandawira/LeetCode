package array

func containsDuplicate(nums []int) bool {
	prevs := make(map[int]struct{}, len(nums))

	for _, num := range nums {

		if _, ok := prevs[num]; ok {
			return true
		}

		prevs[num] = struct{}{}
	}

	return false
}
