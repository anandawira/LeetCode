package dp

func coinChange(coins []int, amount int) int {
	minCoins := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		min := amount + 1
		for _, denom := range coins {
			if denom > i {
				continue
			}

			if minCoins[i-denom] == -1 {
				continue
			}

			cur := 1 + minCoins[i-denom]
			if cur < min {
				min = cur
			}
		}

		if min != amount+1 {
			minCoins[i] = min
		} else {
			minCoins[i] = -1
		}
	}
	return minCoins[amount]
}