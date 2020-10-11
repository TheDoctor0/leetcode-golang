package solutions

import (
    "sort"
)

func largestDivisibleSubset(nums []int) []int {
    sort.Slice(nums, func (i int, j int) bool {
        return nums[i] > nums[j]
    })

    dictionary := make(map[int][]int)

    return findLargestDivisibleSubset(nums, dictionary, 0)
}

func findLargestDivisibleSubset(nums []int, dictionary map[int][]int, current int) []int {
    if len(nums) == 0 {
        return []int{}
    }

    var result []int

    for i := 0; i < len(nums); i++  {
        var temp []int

        if current% nums[i] == 0 {
            if value, ok := dictionary[nums[i]]; ok {
                temp = value
            } else {
                temp = append([]int{nums[i]}, findLargestDivisibleSubset(nums[i + 1:], dictionary, nums[i])...)
            }

            dictionary[nums[i]] = temp

            if len(temp) > len(result) {
                result = temp
            }
        }
    }

    return result
}
