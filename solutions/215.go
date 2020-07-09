package solutions

func findKthLargest(nums []int, k int) int {
    base := len(nums) / 2 - 1

    for i := base; i >= 0; i-- {
        siftUp(nums, i, len(nums) - 1)
    }

    for i := len(nums) - 1; i >= 0; i-- {
        k--

        if k == 0 {
            return nums[0]
        }

        nums[0], nums[i] = nums[i], nums[0]

        siftUp(nums, 0, i - 1)
    }

    return nums[0]
}

func siftUp(nums []int, i, end int) {
    largest, left, right := i, 2 * i + 1, 2 * i + 2

    if left <= end && nums[largest] < nums[left] {
        largest = left
    }

    if right <= end && nums[largest] < nums[right] {
        largest = right
    }

    if i != largest {
        nums[i], nums[largest] = nums[largest], nums[i]

        siftUp(nums, largest, end)
    }
}
