package solutions

import (
    "math"
)

func increasingTriplet(nums []int) bool {
    if nums == nil || len(nums) < 3 {
        return false
    }

    first, second := nums[0], math.MaxInt16

    for i := 1; i < len(nums); i++ {
        if nums[i] <= first{
            first = nums[i]
        } else if  second != math.MaxInt16 && nums[i] > second {
            return true
        } else {
            second = nums[i]
        }
    }

    return false
}
