package solutions

func searchInsert(nums []int, target int) int {
    if target < nums[0] {
        return 0
    }

    for index, value := range nums {
        if target <= value {
            return index
        }
    }

    return len(nums)
}