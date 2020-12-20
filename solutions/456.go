package solutions

func find132pattern(nums []int) bool {
    stack := make([]int, 0)
    max := -1 << 31

    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] < max {
            return true
        }

        for len(stack) > 0 && stack[len(stack) - 1] < nums[i] {
            max = stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
        }

        stack = append(stack, nums[i])
    }

    return false
}
