package random

func maxProfit(prices []int) int {
	var highestProfit int
	buyPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		curPrice := prices[i]

		if curPrice < buyPrice {
			buyPrice = curPrice
			continue
		}

		if diff := curPrice - buyPrice; diff > highestProfit {
			highestProfit = diff
		}
	}

	return highestProfit
}
