package solutions

import (
    "unicode"
)

func calculate(s string) int {
    stack := make([]int, 0, len(s))
    result, number, sign := 0, 0, 1

    for i := 0; i < len(s); i++ {
        switch s[i] {
        case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
            j := i
            number = 0

            for j < len(s) && unicode.IsDigit(rune(s[j])) == true {
                number = number * 10 + int(s[j] - '0')
                j++
            }

            result += sign * number
            i = j - 1
        case '+':
            sign = 1
        case '-':
            sign = -1
        case '(':
            stack = append(stack, result)
            stack = append(stack, sign)

            result = 0
            sign = 1
        case ')':
            sign = stack[len(stack) - 1]
            previousLevelResult := stack[len(stack) - 2]
            stack = stack[: len(stack) - 2]

            result = previousLevelResult + sign * result
        }
    }

    return result
}
