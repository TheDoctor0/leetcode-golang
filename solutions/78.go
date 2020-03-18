package solutions

func subsets(nums []int) [][]int {
    if len(nums) == 0 {
        return [][]int{{}}
    }

    result := make([][]int, 0)

    backtrackSubsets(nums, []int{}, &result)

    return result
}

func backtrackSubsets(nums []int, path []int, result *[][]int) {
    *result = append(*result, path)

    for i := 0; i < len(nums); i++ {
        nextPath := append([]int{}, path...)

        backtrackSubsets(nums[i + 1:], append(nextPath, nums[i]), result)
    }
}
