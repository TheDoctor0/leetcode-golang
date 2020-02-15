package solutions

func permute(nums []int) [][]int {
    if len(nums) == 1 {
        return [][]int{nums}
    }

    first, remaining := nums[0], permute(nums[1:])
    length := len(remaining[0]) + 1
    result := make([][]int, len(remaining) * length)

    for i, permutation := range remaining {
        for j := 0; j < length; j++ {
            temp := append([]int{}, permutation[:j]...)
            temp = append(temp, first)
            temp = append(temp, permutation[j:]...)

            result[i * length + j] = temp
        }
    }

    return result
}
