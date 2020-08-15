package solutions

func firstBadVersion(n int) int {
    i, j := 1, n

    for i + 1 < j {
        middle := (i + j) / 2
        if isBadVersion(middle) {
            j = middle
        } else {
            i = middle
        }
    }

    if isBadVersion(i) {
        return i
    }

    return j
}
