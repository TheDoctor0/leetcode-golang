package solutions

func longestValidParentheses(s string) int {
    var stack = make([]int, len(s) + 1)
    currentStart := -1
    currentIndex := 0
    result := 0

    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            stack[currentIndex] = currentStart
            currentStart = i
            currentIndex++
        } else if currentIndex == 0 {
            currentStart = i
        } else {
            currentIndex--
            currentStart = stack[currentIndex]

            if i - currentStart > result {
                result = i - currentStart
            }
        }
    }

    return result
}
