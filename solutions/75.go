package solutions

func sortColors(nums []int) {
    low, mid, high := 0, 0, len(nums) - 1

    for mid <= high {
        if nums[mid] == 0 {
            swapIndexes(nums, low, mid)

            low++
            mid++
        } else if nums[mid] == 2 {
            swapIndexes(nums, mid, high)

            high--

        } else if nums[mid] == 1 {
            mid++
        }
    }
}

func swapIndexes(array []int, i, j int) {
    array[i], array[j] = array[j], array[i]
}
