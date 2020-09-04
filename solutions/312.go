package solutions

func maxCoins(nums []int) int {
    length := len(nums)
    nums = append([]int{1}, append(nums, 1)...)
    dp := make([][]int, length + 2)

    for i := range dp {
        dp[i] = make([]int, length + 2)
    }

    for l := 1; l <= length; l++ {
        for i := 1; i <= length - l + 1; i++ {
            j := i + l - 1

            for k := i; k <= j; k++ {
                dp[i][j] = max(dp[i][j], dp[i][k - 1] + nums[i - 1] * nums[k] * nums[j + 1] + dp[k + 1][j])
            }
        }
    }

    return dp[1][length]
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
