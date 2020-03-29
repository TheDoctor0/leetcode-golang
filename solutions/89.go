package solutions

func grayCode(n int) []int {
    if n < 0 {
        return nil
    }

    result := []int{0}

    if n == 0 {
        return result
    }

    for i := 0; i < n; i++ {
        for j := len(result) - 1; j >= 0; j-- {
            result = append(result, result[j] | (1 << uint(i)))
        }
    }

    return result
}
