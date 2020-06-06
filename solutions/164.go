package solutions

import (
    "math"
)

func maximumGap(nums []int) int {
    if len(nums) < 2 {
        return 0
    }

    min, max := nums[0], nums[0]

    for _, value := range nums {
        if min > value {
            min = value
        }

        if max < value {
            max = value
        }
    }

    bucketSize := int(math.Ceil(float64(max - min) / float64(len(nums) - 1)))
    bucketsMin, bucketsMax := make([]int, len(nums) - 1), make([]int, len(nums) - 1)

    for i := 0; i < len(nums) - 1; i++ {
        bucketsMin[i], bucketsMax[i] = math.MaxInt32, math.MinInt32
    }

    for _, value := range nums {
        if value == min || value == max {
            continue
        }

        index := (value - min) / bucketSize

        if value < bucketsMin[index] {
            bucketsMin[index] = value
        }

        if value > bucketsMax[index] {
            bucketsMax[index] = value
        }
    }

    maxGap, previous := math.MinInt32, min

    for i := 0; i < len(nums) - 1; i++ {
        if bucketsMin[i] == math.MaxInt32 && bucketsMax[i] == math.MinInt32 {
            continue
        }

        if bucketsMin[i] - previous > maxGap {
            maxGap = bucketsMin[i] - previous
        }

        previous = bucketsMax[i]
    }

    if max - previous > maxGap {
        maxGap = max - previous
    }

    return maxGap
}