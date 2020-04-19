package solutions

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}

func max(a int, b int) int {
    if a > b {
        return a
    }

    return b
}

func abs(number int) int {
    if number < 0 {
        return -number
    }

    return number
}