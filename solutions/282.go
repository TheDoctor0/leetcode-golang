package solutions

import (
    "strings"
)

func addOperators(num string, target int) []string {
    result := make([]string, 0)

    if len(num) == 0 {
        return result
    }

    path, digits, n := make([]rune, len(num) * 2 - 1), []rune(num), 0

    for i := 0; i < len(digits); i++ {
        n = n * 10 + int(digits[i]) - '0'
        path[i] = digits[i]

        if i < len(path) {
            path[i + 1] = ' '
        }

        result = includeOperators(result, path, i + 1, 0, n, digits, i + 1, target)

        if n == 0 {
            break
        }
    }

    return result
}

func includeOperators(result []string, path []rune, length int, left int, current int, digits []rune, index int, target int) []string {
    if index == len(digits) {
        if left + current == target {
            add := string(path)
            add = strings.Replace(add, " ", "", -1)

            result = append(result, add)
        }

        return result
    }

    j, n := length + 1, 0

    for i := index; i < len(digits); i++ {
        n = n * 10 + int(digits[i]) - '0'
        path[j] = digits[i]
        j++

        if j < len(path) {
            path[j] = ' '
        }

        path[length] = '+'
        result = includeOperators(result, path, j, left + current, n, digits, i + 1, target)

        path[length] = '-'
        result = includeOperators(result, path, j, left + current, -n, digits, i + 1, target)

        path[length] = '*'
        result = includeOperators(result, path, j, left, current * n, digits, i + 1, target)

        if digits[index] == '0' {
            break
        }
    }

    return result
}
