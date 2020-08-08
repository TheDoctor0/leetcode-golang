package solutions

func singleNumber(nums []int) []int {
    xor := 0

    for _, value := range nums {
        xor = xor ^ value
    }

    result, lastBit := make([]int, 2), xor & (-xor)

    for _, value := range nums {
        if value & lastBit == 0 {
            result[0] ^= value
        } else {
            result[1] ^= value
        }
    }

    return result
}
