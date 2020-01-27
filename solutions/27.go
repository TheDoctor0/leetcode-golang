package solutions

func removeElement(nums []int, val int) int {
    if nums == nil || len(nums) == 0 {
        return 0
    }

    index := 0

    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[index] = nums[i]
            index++
        }
    }

    return index
}
