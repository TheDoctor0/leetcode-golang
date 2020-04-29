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

func minInArray(values []int) int {
    min := values[0]

    for i := 1; i < len(values); i++ {
        if values[i] < min {
            min = values[i]
        }
    }

    return min
}