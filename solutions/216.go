package solutions

func combinationSum3(k int, n int) [][]int {
    var result [][]int
    var current []int

    searchCombination(k, n, 1, current, &result)

    return result
}

func searchCombination(k, n, depth int, current []int, result *[][]int) {
    if k == 0 && n == 0 {
        *result = append(*result, append([]int{}, current...))

        return
    }

    if k < 0 || n < 0 {
        return
    }

    for i := depth; i <= 9; i++ {
        searchCombination(k - 1, n - i, i + 1, append(current, i), result)
    }
}
