package solutions

func removeDuplicates(nums []int) int {
    if len(nums) <= 2 {
        return len(nums)
    }

    result := 2

    for i := 2; i < len(nums); i++ {
        if nums[result - 1] != nums[result - 2] || nums[result - 1] != nums[i] {
            nums[result] = nums[i]

            result++
        }
    }

    return result
}
