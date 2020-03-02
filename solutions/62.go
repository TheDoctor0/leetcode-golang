package solutions

func uniquePaths(m int, n int) int {
    list := make([]int, m)

    for i := 0; i < m; i++ {
        list[i] = 1
    }

    for j := 1; j < n; j++ {
        for i := 1; i < m; i++ {
            list[i] = list[i] + list[i - 1]
        }
    }

    return list[m - 1]
}
