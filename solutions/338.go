package solutions

func countBits(num int) []int {
    result := make([]int, num + 1)

    for i := 1; i <= num; i++ {
        result[i] = result[i & (i - 1)] + 1
    }

    return result
}
