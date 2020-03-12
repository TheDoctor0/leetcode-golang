package solutions

func minDistance(word1 string, word2 string) int {
    if word1 == word2 {
        return 0
    }

    lengthFirst, lengthSecond := len(word1), len(word2)
    dp := make([][]int, 2)

    for i := 0; i < 2; i++ {
        dp[i] = make([]int, lengthSecond + 1)
    }

    for i := 0; i < lengthFirst + 1; i++ {
        for j := 0; j < lengthSecond + 1; j++ {
            if i == 0 {
                dp[0][j] = j
            } else if j == 0 {
                dp[1][0] = i
            } else if word1[i - 1] == word2[j - 1] {
                dp[1][j] = dp[0][j - 1]
            } else {
                dp[1][j] = min(min(dp[0][j - 1] + 1, dp[1][j - 1] + 1), dp[0][j] + 1)
            }
        }

        if i > 0 {
            for k := 0; k < lengthSecond + 1; k++ {
                dp[0][k] = dp[1][k]
            }
        }
    }

    if lengthFirst == 0 {
        return dp[0][lengthSecond]
    }

    return dp[1][lengthSecond]
}
