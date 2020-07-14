package solutions

import (
    "math"
)

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
    if k <= 0 || t < 0 {
        return false
    }

    bucket := make(map[int64]int64)
    numbers := make([]int64, len(nums))

    for i, number := range nums {
        numbers[i] = int64(number) - int64(math.MinInt32)
    }

    for i, number := range numbers {
        removeIndex := i - (k + 1)

        if removeIndex >= 0 {
            currentIndex := numbers[removeIndex] / (int64(t) + 1)

            delete(bucket, currentIndex)
        }

        currentIndex := number / (int64(t) + 1)

        if _, ok := bucket[currentIndex]; ok {
            return true
        }

        if neighborNumber, ok := bucket[currentIndex + 1]; ok && neighborNumber- number <= int64(t) {
            return true
        }

        if neighborNum, ok := bucket[currentIndex - 1]; ok && number - neighborNum <= int64(t) {
            return true
        }

        bucket[currentIndex] = number
    }

    return false
}
