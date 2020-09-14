package solutions

func countRangeSum(nums []int, lower int, upper int) int {
    length := len(nums)
    sum := make([]int64, length + 1)

    for i := 1; i <= length; i++ {
        sum[i] = sum[i - 1] + int64(nums[i - 1])
    }

    return mergeRangeSum(sum, 0, length, lower, upper)
}

func mergeRangeSum(sum []int64, left, right, lower, upper int) int {
    if left >= right {
        return 0
    }

    middle := (left + right) / 2
    m, n := middle + 1, middle + 1
    count := mergeRangeSum(sum, left, middle, lower, upper) + mergeRangeSum(sum, middle + 1, right, lower, upper)

    for i := left; i <= middle; i++ {
        for m <= right && sum[m] - sum[i] < int64(lower) {
            m++
        }

        for n <= right && sum[n] - sum[i] <= int64(upper) {
            n++
        }

        count += n - m
    }

    leftPart := append([]int64{}, sum[left: middle + 1]...)
    rightPart := append([]int64{}, sum[middle + 1: right+ 1]...)

    i := left
    left, right = 0, 0

    for left < len(leftPart) && right < len(rightPart) {
        if leftPart[left] < rightPart[right] {
            sum[i] = leftPart[left]
            left++
        } else {
            sum[i] = rightPart[right]
            right++
        }

        i++
    }

    for right < len(rightPart) {
        sum[i] = rightPart[right]
        right++
        i++
    }

    for left < len(leftPart) {
        sum[i] = leftPart[left]
        left++
        i++
    }

    return count
}
