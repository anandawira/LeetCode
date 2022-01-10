package main

func ArrayAddition(arr []int) string {
	maxVal := arr[0]
	maxIdx := 0

	for index, val := range arr {
		if val > maxVal {
			maxVal = val
			maxIdx = index
		}
	}

	comArray := append(arr[:maxIdx], arr[maxIdx+1:]...)
	result := dfs(&comArray, 0, 0, maxVal)
	if result {
		return "true"
	} else {
		return "false"
	}
}

func dfs(arr *[]int, index int, total int, target int) bool {
	// Check if out of bound
	if index > len(*arr)-1 {
		return false
	}

	// Check if current total equals to target
	if total == target {
		return true
	}

	// Check if number more than target
	if total > target {
		return false
	}

	// without current index
	result1 := dfs(arr, index+1, total, target)

	// with current index
	result2 := dfs(arr, index+1, total+(*arr)[index], target)

	return result1 || result2
}

func main() {
	ArrayAddition([]int{3, 5, -1, 8, 12})

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	// fmt.Println(ArrayAddition(readline()))

}
