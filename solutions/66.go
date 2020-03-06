package solutions

func plusOne(digits []int) []int {
    for i := len(digits); i > 0; i-- {
        if digits[i - 1] < 9 {
            digits[i - 1] += 1

            return digits
        }

        digits[i - 1] = 0

        if i == 1 {
            digits = append([]int{1}, digits...)
        }
    }

    return digits
}
