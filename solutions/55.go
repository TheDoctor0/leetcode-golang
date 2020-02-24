package solutions

func canJump(nums []int) bool {
    previous := len(nums) - 1

    for i := previous - 1; i >= 0; i-- {
        if nums[i] >= previous - i {
            previous = i
        }
    }

    return previous == 0
}
