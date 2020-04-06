package solutions

func isInterleave(s1, s2, s3 string) bool {
    if len(s3) != len(s1) + len(s2) {
        return false
    }

    var dp [][]bool

    for i := 0; i < len(s3) + 1; i++ {
        dp = append(dp, make([]bool, len(s1) + 1))
    }

    dp[0][0] = true

    for i := 1; i <= len(s3); i++ {
        for j := 0; j <= i; j++ {
            k := i - j

            if j > len(s1) || k > len(s2) {
                continue
            }

            dp[i][j] = (j-1 >= 0 && dp[i-1][j-1] && s3[i-1] == s1[j-1]) || (k-1 >= 0 && dp[i-1][j] && s3[i-1] == s2[k-1])
        }
    }

    return dp[len(s3)][len(s1)]
}
