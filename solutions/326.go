package solutions

func isPowerOfThree(n int) bool {
    for ; n > 1 && n % 3 == 0; n /= 3 {}

    return n == 1
}
