package solutions

func rangeBitwiseAnd(m int, n int) int {
    for n > m && n != 0 {
        n = n & (n - 1)
    }

    return n
}
