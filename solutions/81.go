package solutions

func search(nums []int, target int) bool {
    left, right := 0, len(nums) - 1

    for left <= right {
        middle := (left + right) / 2

        if nums[middle] == target || target == nums[left] || target == nums[right] {
            return true
        }

        if nums[left] < nums[middle] {
            if target < nums[middle] && target >= nums[left] {
                right = middle
            } else {
                left = middle + 1
            }
        } else if nums[left] > nums[middle] {
            if target > nums[middle] && target <= nums[right] {
                left = middle + 1
            } else {
                right = middle
            }
        } else {
            left++
        }
    }

    return false
}