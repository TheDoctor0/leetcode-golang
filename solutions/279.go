package solutions

import (
    "math"
)

func numSquares(n int) int {
    dp := make([]int, n + 1)
    dp[0], dp[1] = 0, 1

    for i := 2; i <= n; i++ {
        dp[i] = math.MaxInt32

        for j := 1; i - j * j >= 0; j++ {
            dp[i] = min(dp[i], dp[i - j * j] + 1)
        }
    }

    return dp[n]
}
