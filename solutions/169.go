package solutions

func majorityElement(nums []int) int {
    count, result := 0, 0

    for i := 0; i < len(nums); i++ {
        if count == 0 {
            result = nums[i]
            count++
        } else if nums[i] == result {
            count++
        } else {
            count--
        }
    }

    return result
}
