package solutions

func twoSum(numbers []int, target int) []int {
    left, right := 0, len(numbers) - 1

    for {
        if numbers[left] + numbers[right] > target {
            right--
        } else if numbers[left] + numbers[right] < target {
            left++
        } else {
            return []int{left + 1, right + 1}
        }
    }
}
