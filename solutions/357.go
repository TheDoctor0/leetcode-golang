package solutions

func countNumbersWithUniqueDigits(n int) int {
    result := 1

    if n > 10 {
        n = 10
    }

    for i := 1; i <= n; i++ {
        j, base, end := i, 9, 9

        for j > 1 {
            base *= end
            end--
            j--
        }

        result += base
    }

    return result
}
