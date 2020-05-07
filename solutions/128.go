package solutions

func longestConsecutive(nums []int) int {
    temp := map[int]int{}
    result := 0

    for _, number := range nums {
        if _, ok := temp[number]; !ok {
            left := temp[number - 1]
            right := temp[number + 1]

            temp[number] = 1 + left + right

            if left > 0 {
                temp[number-left] = temp[number]
            }

            if right > 0 {
                temp[number+right] = temp[number]
            }

            if temp[number] > result {
                result = temp[number]
            }
        }
    }

    return result
}
