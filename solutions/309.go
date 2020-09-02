package solutions

func maxProfit(prices []int) int {
    buy, sell := -int(^uint(0) >> 1) - 1, -int(^uint(0) >> 1) - 1
    nothing := 0

    for i := 0; i < len(prices); i++ {
        temp := nothing

        nothing = max(nothing, sell)
        buy = max(buy, temp - prices[i])

        sell = buy + prices[i]
    }

    return max(nothing, sell)
}
