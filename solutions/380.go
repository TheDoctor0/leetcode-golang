package solutions

import (
    "math/rand"
)

type RandomizedSet struct {
    nums          []int
    numberToIndex map[int]int
}

func Constructor() RandomizedSet {
    return RandomizedSet{[]int{}, make(map[int]int)}
}

func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.numberToIndex[val]; ok {
        return false
    }

    this.nums = append(this.nums, val)
    this.numberToIndex[val] = len(this.nums) - 1

    return true
}

func (this *RandomizedSet) Remove(val int) bool {
    index, ok := this.numberToIndex[val]

    if !ok {
        return false
    }

    if len(this.nums) < 2 {
        this.nums = this.nums[0: 0]
    } else {
        swap := this.nums[len(this.nums) - 1]

        this.numberToIndex[swap] = index
        this.nums[index] = swap
        this.nums = this.nums[0: len(this.nums) - 1]
    }

    delete(this.numberToIndex, val)

    return true
}

func (this *RandomizedSet) GetRandom() int {
    return this.nums[rand.Intn(len(this.numberToIndex))]
}
