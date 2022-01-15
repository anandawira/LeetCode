package tokopedia

func maxInversions(arr []int32) int64 {
	// Write your code here
	var result int64

	for index, val := range arr {
		// Find smaller number on right
		var small int64
		for i := index + 1; i < len(arr); i++ {
			if arr[i] < val {
				small++
			}
		}

		// Find bigger number on left
		var big int64
		for i := 0; i < index; i++ {
			if arr[i] > val {
				big++
			}
		}

		result += small * big
	}

	return result
}
