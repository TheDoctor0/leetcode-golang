package solutions

func minSubArrayLen(s int, nums []int) int {
    var i, j, sum int

    min := 1<<31 - 1

    for j < len(nums) {
        sum += nums[j]
        j++

        for sum >= s {
            if min > j - i {
                min = j - i
            }

            sum -= nums[i]
            i++
        }
    }

    if min == 1<<31 - 1 {
        return 0
    }

    return min
}
