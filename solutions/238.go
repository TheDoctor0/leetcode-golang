package solutions

func productExceptSelf(nums []int) []int {
    forward := make([]int, len(nums))
    backward := make([]int, len(nums))
    result := make([]int, len(nums))

    for i := range nums {
        forward[i], backward[i] = 1, 1
    }

    for i := 1; i < len(nums); i++ {
        forward[i] = forward[i - 1] * nums[i - 1]
    }

    for i := len(nums) - 2; i >= 0; i-- {
        backward[i] = backward[i + 1] * nums[i + 1]
    }

    for i := range nums {
        result[i] = forward[i] * backward[i]
    }

    return result
}
