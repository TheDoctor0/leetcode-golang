package solutions

func numTrees(n int) int {
    result := make([]int, n + 1)
    result[0] = 1

    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            result[i] += result[j - 1] * result[i - j]
        }
    }

    return result[n]
}
