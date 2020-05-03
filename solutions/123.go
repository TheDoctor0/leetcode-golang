package solutions

import (
    "math"
)

func maxProfit(prices []int) int {
    firstBuy, secondBuy := math.MinInt32, math.MinInt32
    firstSell, secondSell := 0, 0

    for _, price := range prices {
        secondSell = max(secondSell, secondBuy + price)
        secondBuy = max(secondBuy, firstSell - price)
        firstSell = max(firstSell, firstBuy + price)
        firstBuy = max(firstBuy, -price)
    }

    return secondSell
}
