package solutions

func hammingWeight(num uint32) int {
    result := 0

    for i := 0; i < 32; i++ {
        if (num & (1 << i)) >> i == 1 {
            result++
        }
    }

    return result
}
