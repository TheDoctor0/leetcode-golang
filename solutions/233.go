package solutions

func countDigitOne(n int) int {
    if n <= 0 {
        return 0
    }

    if n <= 9 {
        return 1
    }

    count := count(n)
    current := count

    if n < 2 * count && n >= count {
        current = n - count + 1
    }

    return n / count * countDigitOne(count - 1) + current + countDigitOne(n % count)
}

func count(n int) int {
    result := 1

    for n > 0 {
        result *= 10
        n /= 10
    }

    return result / 10
}
