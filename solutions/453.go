package solutions

import (
    "math"
)

func minMoves(nums []int) int {
    min := math.MaxInt64

    for _, value := range nums {
        if value < min {
            min = value
        }
    }

    result := 0

    for _, value := range nums {
        result += value - min
    }

    return result
}
