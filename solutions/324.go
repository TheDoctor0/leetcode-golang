package solutions

import (
    "sort"
)

func wiggleSort(nums []int) {
    if len(nums) <= 1 {
        return
    }

    sort.Ints(nums)

    middle := len(nums) / 2 + len(nums) % 2

    var result []int

    for i := 0; i < middle; i++ {
        result = append(result, nums[middle - i - 1])

        if middle + i < len(nums){
            result = append(result, nums[len(nums) - i - 1])
        }
    }

    for i := 0; i < len(result); i++ {
        nums[i] = result[i]
    }
}
