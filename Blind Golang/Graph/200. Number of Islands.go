package graph

func drownIsland(i, j int, grid [][]byte) {
	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'

	if i != 0 {
		drownIsland(i-1, j, grid)
	}
	if i != len(grid)-1 {
		drownIsland(i+1, j, grid)
	}
	if j != 0 {
		drownIsland(i, j-1, grid)
	}
	if j != len(grid[0])-1 {
		drownIsland(i, j+1, grid)
	}
}

func numIslands(grid [][]byte) int {
	var result int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				result++
				drownIsland(i, j, grid)
			}
		}
	}

	return result
}
