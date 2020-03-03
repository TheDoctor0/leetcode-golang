package solutions

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    rows, columns := len(obstacleGrid), len(obstacleGrid[0])
    result := make([][]int, rows)

    if obstacleGrid[0][0] == 1 {
        return 0
    }

    for i := range result {
        result[i] = make([]int, columns)
    }

    result[0][0] = 1

    for i := 1; i < rows; i++ {
        if obstacleGrid[i][0] == 1 || result[i - 1][0] == 0 {
            result[i][0] = 0
        } else {
            result[i][0] = 1
        }
    }

    for i := 1; i < columns; i++ {
        if obstacleGrid[0][i] == 1 || result[0][i - 1] == 0 {
            result[0][i] = 0
        } else {
            result[0][i] = 1
        }
    }

    for row := 1; row < rows; row++ {
        for column := 1; column < columns; column++ {
            if obstacleGrid[row][column] != 1 {
                result[row][column] = result[row - 1][column] + result[row][column - 1]
            } else {
                result[row][column] = 0
            }
        }
    }

    return result[rows - 1][columns - 1]
}
