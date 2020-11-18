package solutions

func numberOfArithmeticSlices(A []int) int {
    dp, result := 0, 0

    for i := 2; i < len(A); i++ {
        if A[i] - A[i - 1] == A[i - 1] - A[i - 2] {
            dp = dp + 1
            result += dp
        } else {
            dp = 0
        }
    }

    return result
}
