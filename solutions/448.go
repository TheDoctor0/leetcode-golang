package solutions

func findDisappearedNumbers(nums []int) []int {
    for i := 0; i < len(nums); i++ {
        for nums[i] != i + 1 {
            current := nums[i]

            if nums[current - 1] == current {
                break
            } else {
                nums[i], nums[current - 1] = nums[current - 1], nums[i]
            }
        }
    }

    var result []int

    for i, value := range nums {
        if value != i + 1 {
            result = append(result, i + 1)
        }
    }

    return result
}
