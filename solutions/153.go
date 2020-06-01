package solutions

func findMin(nums []int) int {
    left, right := 0, len(nums) - 1

    for left < right {
        middle := (right-left) / 2 + left

        if nums[middle] < nums[0] {
            right = middle
        } else if nums[middle] > nums[right] {
            left = middle + 1
        } else {
            return nums[left]
        }
    }

    return nums[left]
}
