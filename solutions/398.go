package solutions

import (
    "math/rand"
)

type Solution struct {
    list []int
}

func Constructor(nums []int) Solution {
    return Solution{list: nums}
}

func (this *Solution) Pick(target int) int {
    result, count := 0, 0

    for index, value := range this.list {
        if value == target {
            if rand.Intn(count + 1) == 0 {
                result = index
            }

            count++
        }
    }

    return result
}
