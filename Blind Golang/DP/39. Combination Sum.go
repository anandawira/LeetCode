package dp

func combinationSum(candidates []int, target int) [][]int {

	dp := make([][][]int, target+1)
	for i := 0; i <= target; i++ {
		dp[i] = [][]int{}
	}

	for i := 1; i <= target; i++ {
		for j := 0; j < len(candidates); j++ {
			if candidates[j] > i {
				continue
			}

			if candidates[j] == i {
				newCom := []int{candidates[j]}
				dp[i] = append(dp[i], newCom)
			}

			for _, com := range dp[i-candidates[j]] {
				if com[0] < candidates[j] {
					continue
				}

				newCom := append([]int{candidates[j]}, com...)
				dp[i] = append(dp[i], newCom)
			}
		}
	}

	return dp[target]
}
