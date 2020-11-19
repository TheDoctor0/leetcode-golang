package solutions

import (
    "math"
)

func thirdMax(nums []int) int {
    max, second, third := math.MinInt64, math.MinInt64, math.MinInt64

    for _, value := range nums {
        if value == max || value == second || value == third {
            continue
        }

        switch {
        case value > max:
            max, second, third = value, max, second
        case value > second:
            second, third = value, second
        case value > third:
            third = value
        }
    }

    if third == math.MinInt64 {
        return max
    }

    return third
}
