package solutions

func findMaximumXOR(nums []int) int {
    result := 0

    for i := 31; i >= 0; i-- {
        result <<= 1
        pre := make(map[int]bool)

        for _, number := range nums {
            pre[number >> i] = true
        }

        for p := range pre {
            if pre[result ^ 1 ^ p] {
                result += 1
                break
            }
        }
    }

    return result
}
