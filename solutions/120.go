package solutions

func minimumTotal(triangle [][]int) int {
    for i := 0; i < len(triangle); i++ {
        for j := 0; j <= i; j++ {
            if i == 0 && j == 0 {
                continue
            }

            if j == 0 {
                triangle[i][j] += triangle[i - 1][j]
            } else if j == i {
                triangle[i][j] += triangle[i - 1][j - 1]
            } else {
                triangle[i][j] += min(triangle[i - 1][j], triangle[i - 1][j - 1])
            }
        }
    }

    return minInArray(triangle[len(triangle) - 1])
}
