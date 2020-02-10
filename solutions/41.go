package solutions

func firstMissingPositive(nums []int) int {
    length := len(nums)

    for i := 0; i < length; {
        if nums[i] > 0 && nums[i] <= length && nums[i] != nums[nums[i] - 1] {
            nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
        } else {
            i++
        }
    }

    for i := 0; i < length; i++ {
        if nums[i] != i + 1 {
            return i + 1
        }
    }

    return length + 1
}