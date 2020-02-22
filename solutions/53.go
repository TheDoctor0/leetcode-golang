package solutions

func maxSubArray(nums []int) int {
    max, sum := nums[0],nums[0]

    for _, value := range nums[1:] {
        if sum < 0 {
            sum = value
        } else {
            sum += value
        }

        if max < sum {
            max = sum
        }
    }

    return max
}
