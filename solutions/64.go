package solutions

func minPathSum(grid [][]int) int {
    rowLength := len(grid)
    columnLength := len(grid[0])

    for i := 1; i < columnLength; i++ {
        grid[0][i] += grid[0][i - 1]
    }

    for i := 1; i < rowLength; i++ {
        grid[i][0] += grid[i - 1][0]
    }

    for i := 1; i < rowLength; i++ {
        for j := 1; j < columnLength; j++ {
            left := grid[i][j - 1] + grid[i][j]
            up := grid[i - 1][j] + grid[i][j]

            if left > up {
                grid[i][j] = up
            } else {
                grid[i][j] = left
            }
        }
    }

    return grid[rowLength - 1][columnLength - 1]
}
