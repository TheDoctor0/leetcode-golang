package solutions

func rob(nums []int) int {
    rob, dontRob := 0, 0

    for _, number := range nums {
        rob, dontRob = number + dontRob, max(dontRob, rob)
    }

    return max(dontRob, rob)
}
