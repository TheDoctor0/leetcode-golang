package solutions

func combine(n int, k int) [][]int {
    var result [][]int

    for i := 1; i <= n-k + 1; i++ {
        backtrackCombinations(&result, []int{i}, n, k)
    }

    return result
}

func backtrackCombinations(result *[][]int, previous []int, n, k int) {
    if len(previous) == k {
        *result = append(*result, previous)

        return
    }

    if previous[len(previous) - 1] == n {
        return
    }

    for v := previous[len(previous) - 1]; v < n && n - v >= k - len(previous); v++ {
        backtrackCombinations(result, append(append([]int{}, previous...), v + 1), n, k)
    }
}
