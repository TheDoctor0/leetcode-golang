package solutions

func findMin(nums []int) int {
    left, right := 0, len(nums) - 1

    for left < right {
        middle := (right + left) / 2
        if nums[middle] < nums[right] {
            right = middle
        } else if nums[middle] > nums[right] {
            left = middle + 1
        } else if nums[middle] != nums[left] {
            right = middle
        } else {
            return searchMin(nums[left : right + 1])
        }
    }

    return nums[left]
}

func searchMin(nums []int) int {
    result := nums[0]

    for i := 1; i < len(nums); i++ {
        if nums[i] < result {
            result = nums[i]
        }
    }

    return result
}
