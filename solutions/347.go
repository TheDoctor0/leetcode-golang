package solutions

import (
    "sort"
)

func topKFrequent(nums []int, k int) []int {
    dictionary := make(map[int]int)

    for _, value := range nums {
        dictionary[value]++
    }

    var keys []int

    for key := range dictionary {
        keys = append(keys, key)
    }

    sort.Slice(keys, func (i int, j int) bool {
        return dictionary[keys[i]] > dictionary[keys[j]]
    })

    return keys[:k]
}
