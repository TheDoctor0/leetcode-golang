package solutions

import (
    "math"
)

func findDuplicates(nums []int) []int {
    result := make([]int, 0)

    for i := range nums {
        index := int(math.Abs(float64(nums[i])) - 1)

        if nums[index] < 0 {
            result = append(result, index + 1)
        }

        nums[index] = -nums[index]
    }

    return result
}
