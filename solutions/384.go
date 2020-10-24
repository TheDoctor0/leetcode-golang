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

func (this *Solution) Reset() []int {
    return this.list
}

func (this *Solution) Shuffle() []int {
    shuffled := make([]int, len(this.list))
    copy(shuffled, this.list)
    length := len(shuffled)

    for i := 0; i < length; i++ {
        j := rand.Intn(length - i)
        shuffled[i], shuffled[j + i] = shuffled[j + i], shuffled[i]
    }

    return shuffled
}
