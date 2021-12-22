package array

func maxProfit(prices []int) int {
	var highestProfit int
	buyPrice := &prices[0]

	for i := 1; i < len(prices); i++ {
		if diff := prices[i] - *buyPrice; diff > highestProfit {
			highestProfit = diff
		}

		if prices[i] < *buyPrice {
			buyPrice = &prices[i]
		}
	}

	return highestProfit
}
