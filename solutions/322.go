package solutions

import (
    "math"
)

func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)

    return coinChangeHelper(coins, amount, dp)
}

func coinChangeHelper(coins []int, amount int, dp []int) int {
    if amount == 0 || dp[amount] != 0 {
        return dp[amount]
    }

    result := math.MaxInt64

    for _, value := range coins {
        if value == amount {
            result = 1
        } else if value < amount {
            current := 1 + coinChangeHelper(coins, amount -value, dp)

            if current != 0 && current < result {
                result = current
            }
        }
    }

    if result == math.MaxInt64 {
        dp[amount] = -1

        return -1
    }

    dp[amount] = result

    return result
}
