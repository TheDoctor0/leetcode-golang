package solutions

import (
    "sort"
)

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)

    sum := nums[0] + nums[1] + nums[2]
    result := sum
    difference := abs(result - target)
    length := len(nums)

    for i := 0; i < length - 2; i++ {
        for j, k := i + 1, length - 1; j < k; {
            sum = nums[i] + nums[j] + nums[k]

            if sum - target < 0 {
                j++
            } else {
                k--
            }

            if difference > abs(sum - target) {
                difference = abs(sum - target)
                result = sum
            }
        }
    }

    return result
}
