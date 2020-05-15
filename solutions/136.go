package solutions

func singleNumber(nums []int) int {
    var result int

    for number := range nums {
        result ^= nums[number]
    }

    return result
}
