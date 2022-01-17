package random

func minimumTotal(triangle [][]int) int {
	// Loop from second lowest layer to top.
	for i := len(triangle) - 2; i >= 0; i-- {
		// Iterate each element of row
		for j := 0; j <= i; j++ {
			// Add each element of row with minimum of 2 element from lower row.
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
