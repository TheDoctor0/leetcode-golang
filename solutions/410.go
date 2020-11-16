package solutions

func splitArray(nums []int, m int) int {
    left, right := 0, 0

    for _, value := range nums {
        right += value

        if value > left {
            left = value
        }
    }

    for left < right {
        middle := (left + right) / 2
        cut := getCutSize(nums, middle)

        if cut > m {
            left = middle + 1
        } else {
            right = middle
        }
    }

    return left
}

func getCutSize(nums []int, max int) int {
    cutSize, currentValue := 1, 0

    for _, value := range nums {
        if currentValue + value <= max {
            currentValue += value
        } else {
            currentValue = value
            cutSize++
        }
    }

    return cutSize
}
