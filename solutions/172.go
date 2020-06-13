package solutions

func trailingZeroes(n int) int {
    result := 0

    for n != 0 {
        result += n / 5

        n /= 5
    }

    return result
}
