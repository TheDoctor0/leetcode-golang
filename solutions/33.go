package solutions

func search(nums []int, target int) int {
    left := 0
    right := len(nums) - 1

    for left <= right {
        if nums[left] == target {
            return left
        }

        if nums[right] == target {
            return right
        }

        left++
        right--
    }

    return -1
}
