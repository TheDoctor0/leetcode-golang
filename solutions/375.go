package solutions

import (
    "math"
)

func getMoneyAmount(n int) int {
    dp := make([][]int, n + 1)

    for i := 1; i <= n; i ++{
        dp[i] = make([]int, n + 1)
    }

    for i := 1; i < n; i ++{
        dp[i][i + 1] = i
    }

    for k := 2; k < n; k ++{
        for i := 1; i + k <= n; i ++{
            min := math.MaxInt32

            for j := i + 1; j <= i + k - 1; j ++ {
                if current := max(dp[i][j - 1], dp[j + 1][i + k]) + j; current < min {
                    min = current
                }
            }

            dp[i][i + k] = min
        }
    }

    return dp[1][n]
}
