package solutions

import (
    "sort"
)

func subsetsWithDup(nums []int) [][]int {
    result := make([][]int, 0)

    sort.Ints(nums)

    backtrackSubsetsWithDuplicated(nums, []int{}, &result)

    return result
}

func backtrackSubsetsWithDuplicated(nums, path []int, result *[][]int) {
    *result = append(*result, append([]int{}, path...))

    for i := 0; i < len(nums); {
        path = append(path, nums[i])

        backtrackSubsetsWithDuplicated(nums[i + 1:], path, result)

        path = path[:len(path) - 1]

        i++

        for i < len(nums) && nums[i] == nums[i - 1] {
            i++
        }
    }
}
