package solutions

func generateParenthesis(n int) []string {
    result := make([]string, 0)

    generatePairs("(", n - 1, n, &result)

    return result
}

func generatePairs(str string, left, right int, result *[]string) {
    if left > right {
        return
    }

    if left > 0 {
        generatePairs(str + "(", left - 1, right, result)
    }

    if right > 0 {
        generatePairs(str + ")", left, right - 1, result)
    }

    if left == 0 && right == 0 {
        *result = append(*result, str)
    }
}
