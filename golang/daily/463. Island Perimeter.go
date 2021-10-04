package daily

func islandPerimeter(grid [][]int) int {
	res := 0

	for i, l := 0, len(grid); i < l; i++ {
		for j, m := 0, len(grid[0]); j < m; j++ {
			// Check if current cell is 1
			if grid[i][j] == 1 {
				// Add 4 to result as 4 side of square
				res += 4

				// Check if cell on the top is also 1, then subtract 2 from result
				if i > 0 && grid[i-1][j] == 1 {
					res -= 2
				}

				// Check if cell on the left is also 1, then subtract 2 from result
				if j > 0 && grid[i][j-1] == 1 {
					res -= 2
				}
			}
		}
	}

	return res
}
