package solutions

import (
    "math/rand"
)

type Pair struct {
    value int
    index int
}

type RandomizedCollection struct {
    mapToIndex map[int][]int
    nums       []Pair
}

func Constructor() RandomizedCollection {
    return RandomizedCollection {
        mapToIndex: make(map[int][]int),
        nums:       []Pair{},
    }
}

func (this *RandomizedCollection) Insert(val int) bool {
    result := false

    if _, ok := this.mapToIndex[val]; !ok {
        this.mapToIndex[val] = []int{}
        result = true
    }

    this.mapToIndex[val] = append(this.mapToIndex[val], len(this.nums))
    this.nums = append(this.nums, Pair{value: val, index: len(this.mapToIndex[val]) - 1})

    return result
}

func (this *RandomizedCollection) Remove(val int) bool {
    if _, ok := this.mapToIndex[val]; !ok {
        return false
    }

    indices := this.mapToIndex[val]
    replaceIndex := indices[len(indices) - 1]
    last := this.nums[len(this.nums) - 1]

    this.mapToIndex[last.value][last.index] = replaceIndex
    this.nums[replaceIndex] = last
    this.nums = this.nums[:len(this.nums) - 1]
    this.mapToIndex[val] = indices[:len(indices) - 1]

    if len(this.mapToIndex[val]) == 0 {
        delete(this.mapToIndex, val)
    }

    return true
}

func (this *RandomizedCollection) GetRandom() int {
    return this.nums[rand.Intn(len(this.nums))].value
}