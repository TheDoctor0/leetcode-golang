package solutions

func missingNumber(nums []int) int {
    result := len(nums)

    for i, value := range nums {
        result += i - value
    }

    return result
}
