package random

func findNumbers(nums []int) int {
	var result int

	for _, num := range nums {
		digit := getDigits(num)
		if digit%2 == 0 {
			result++
		}
	}

	return result
}

func getDigits(num int) int {
	result := 1

	for num >= 10 {
		result++
		num /= 10
	}

	return result
}
