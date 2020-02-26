package solutions

import (
    "sort"
)

func permuteUnique(nums []int) [][]int {
    if len(nums) == 0 {
        return nil
    }

    sort.Ints(nums)

    result := make([][]int, 0)

    backtrackUnique(nums, nil, &result)

    return result
}

func backtrackUnique(nums []int, previous []int, result *[][]int) {
    if len(nums) == 0 {
        *result = append(*result, append([]int{}, previous...))

        return
    }

    for i := 0; i < len(nums); i++ {
        if i != 0 && nums[i] == nums[i - 1] {
            continue
        }

        current := append(append([]int{}, nums[0:i]...), nums[i + 1:]...)

        backtrackUnique(current, append(previous, nums[i]), result)
    }
}