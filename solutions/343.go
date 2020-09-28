package solutions

func integerBreak(n int) int {
    if n == 2 {
        return 1
    }

    if n == 3 {
        return 2
    }

    product := 1

    for n > 4 {
        product *= 3
        n -= 3
    }

    return product * n
}
