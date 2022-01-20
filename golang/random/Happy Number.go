package random

func isHappy(n int) bool {
	x, y := n, n
	for {
		x = digitSquaredSum(x)
		y = digitSquaredSum(digitSquaredSum(y))
		if x == 1 || y == 1 {
			return true
		}

		if x == y {
			return false
		}
	}
}

func digitSquaredSum(n int) int {
	var sum int
	for n != 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}

	return sum
}
