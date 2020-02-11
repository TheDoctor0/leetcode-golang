package solutions

func trap(h []int) int {
    if len(h) < 3 {
        return 0
    }

    var result, left, right, maxLeft, maxRight int
    right = len(h) - 1

    for left < right {
        if h[left] > maxLeft {
            maxLeft = h[left]
        }

        if h[right] > maxRight {
            maxRight = h[right]
        }

        if maxLeft < maxRight {
            result += maxLeft - h[left]
            left++
        } else {
            result += maxRight - h[right]
            right--
        }
    }

    return result
}
