package solutions

func uniquePaths(m int, n int) int {
    result := make([]int, m)

    for i := 0; i < m; i++ {
        result[i] = 1
    }

    for j := 1; j < n; j++ {
        for i := 1; i < m; i++ {
            result[i] = result[i] + result[i - 1]
        }
    }

    return result[m - 1]
}
