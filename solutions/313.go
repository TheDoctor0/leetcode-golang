package solutions

import (
    "math"
)

func nthSuperUglyNumber(n int, primes []int) int {
    if n <= 1 || len(primes) == 0 {
        return 1
    }

    dp := make([]int, n)
    dp[0] = 1
    numbers := make([]int, len(primes))

    for i := 1; i < n; i++ {
        max := math.MaxInt32

        for index, prime := range primes {
            max = min(max, dp[numbers[index]] * prime)
        }

        dp[i] = max

        for index, prime := range primes {
            if max == dp[numbers[index]] * prime {
                numbers[index]++
            }
        }
    }

    return dp[n - 1]
}
