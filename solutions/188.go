package solutions

func maxProfit(k int, prices []int) int {
    if k >= len(prices) / 2 {
        var profit int

        for i := 1; i < len(prices); i++ {
            if prices[i] > prices[i - 1] {
                profit += prices[i] - prices[i - 1]
            }
        }

        return profit
    }

    profit := make([][]int, k + 1)

    for i := 0; i <= k; i++ {
        profit[i] = make([]int, len(prices))
    }

    for i := 1; i <= k; i++ {
        currentMax := -prices[0]

        for j := 1; j < len(prices); j++ {
            profit[i][j] = max(profit[i][j - 1], prices[j] + currentMax)
            currentMax = max(currentMax, profit[i - 1][j - 1] - prices[j])
        }
    }

    return profit[k][len(prices) - 1]
}
