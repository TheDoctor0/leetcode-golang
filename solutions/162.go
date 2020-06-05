package solutions

func findPeakElement(nums []int) int {
    return searchPeakElement(nums, 0, len(nums) - 1)
}

func searchPeakElement(nums []int, left int, right int) int {
    if left == right {
        return left
    }

    middle := left + (right - left) / 2

    if nums[middle] > nums[middle + 1] {
        return searchPeakElement(nums, left, middle)
    }

    return searchPeakElement(nums, middle + 1, right)
}