package solutions

func lengthOfLIS(nums []int) int {
    tails, size := make([]int, len(nums)), 0

    for _, number := range nums {
        i, j := 0, size

        for i != j {
            middle := (i + j) / 2

            if tails[middle] < number {
                i = middle + 1
            } else {
                j = middle
            }
        }

        tails[i] = number

        if i == size {
            size++
        }
    }

    return size
}
