package solutions

func minPatches(nums []int, n int) int {
    length := len(nums)
    maxValue, result, i := 0, 0, 0

    for maxValue < n {
        if i < length && nums[i] <= maxValue+ 1 {
            maxValue += nums[i]
            i++
        } else {
            maxValue += maxValue + 1
            result++
        }
    }

    return result
}
