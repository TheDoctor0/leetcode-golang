package solutions

import (
    "math"
)

func maxProfit(prices []int) int {
    low := math.MaxInt32
    profit := 0

    for i := 0; i < len(prices); i++ {
        if prices[i] < low {
            low = prices[i]
        } else {
            profit = max(profit, prices[i] - low)
        }
    }

    return profit
}
