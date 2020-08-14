package solutions

func hIndex(citations []int) int {
    length := len(citations)
    left, right := 0, length - 1

    for left <= right {
        middle := left + (right - left) / 2

        if citations[middle] < length-middle {
            left = middle + 1
        } else {
            right = middle - 1
        }
    }

    return length - left
}
