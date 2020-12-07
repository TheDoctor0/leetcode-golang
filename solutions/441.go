package solutions

func arrangeCoins(n int) int {
    n = n * 2
    low, high := 0, n

    for low <= high {
        middle := (high + low) / 2
        value := middle * (middle + 1)

        if value == n {
            return middle
        } else if value < n {
            low = middle + 1
        } else {
            high = middle - 1
        }
    }

    return high
}