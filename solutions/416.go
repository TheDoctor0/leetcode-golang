package solutions

func canPartition(nums []int) bool {
    var sum int

    for _, value := range nums {
        sum += value
    }

    if sum % 2 == 1 {
        return false
    }

    sum = sum / 2
    dp := make([]bool, sum + 1)
    dp[0] = true

    for _, value := range nums {
        for j := sum; j > 0; j-- {
            if j >= value && dp[j - value] {
                dp[j] = dp[j - value]
            }
        }
    }

    return dp[sum]
}
