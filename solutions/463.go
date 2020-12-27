package solutions

func islandPerimeter(grid [][]int) int {
	var lands, neighbours int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}

			lands++

			if i != len(grid) - 1 && grid[i + 1][j] == 1 {
				neighbours++
			}

			if j != len(grid[0]) - 1 && grid[i][j + 1] == 1 {
				neighbours++
			}
		}
	}

	return 4 * lands - 2 * neighbours
}
