package daily

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	isIncreasing := arr[1] > arr[0]

	if !isIncreasing {
		return false
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false
		}

		if isIncreasing {
			if arr[i] < arr[i-1] {
				isIncreasing = false
			}
		} else {
			if arr[i] > arr[i-1] {
				return false
			}
		}
	}

	return !isIncreasing
}
