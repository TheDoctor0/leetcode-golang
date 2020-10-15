package solutions

func guessNumber(n int) int {
    left, right := 1, n

    for left <= right {
        middle := (left + right) / 2
        result := guess(middle)

        if result == 0 {
            return middle
        } else if result == -1 {
            right = middle -1
        } else {
            left = middle +1
        }
    }

    return -1
}
