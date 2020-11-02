package solutions

import (
    "strings"
)

func decodeString(s string) string {
    stackNumber, stackString := make([]int, 0), make([]string, 0)

    var result string
    var number int

    for i := 0; i < len(s); i++ {
        switch {
        case s[i] >= '0' && s[i] <= '9':
            number = 10 * number + int(s[i]) - '0'
        case s[i] == '[':
            stackNumber = append(stackNumber, number)
            stackString = append(stackString, result)
            number = 0
            result = ""
        case s[i] == ']':
            temp := stackString[len(stackString) - 1]
            stackString = stackString[:len(stackString) - 1]
            count := stackNumber[len(stackNumber)  -1]
            stackNumber = stackNumber[:len(stackNumber) - 1]
            result = temp + strings.Repeat(result, count)
        default:
            result += string(s[i])
        }
    }

    return result
}
