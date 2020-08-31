package solutions

import (
    "strconv"
)

func isAdditiveNumber(s string) bool {
    length := len(s)

    for i := 1; i < length - 1; i++ {
        if i > 1 && s[0] == '0' {
            continue
        }

        first, _ := strconv.Atoi(s[: i])

        for j := i + 1; j < length; j++ {
            if j > i + 1 && s[i] == '0' {
                continue
            }

            second, _ := strconv.Atoi(s[i: j])

            if isValidAdditiveNumber(first, second, j, length, s) {
                return true
            }
        }
    }

    return false
}

func isValidAdditiveNumber(first, second, index, length int, s string) bool {
    if index == length {
        return true
    }

    next := strconv.Itoa(first+second)

    if index + len(next) > length {
        return false
    }

    number := 0

    for i := 0; i < len(next); i++ {
        if s[index + i] != next[i] {
            return false
        }

        number = number * 10 + int(s[index + i] - '0')
    }

    first, second = second, number

    return isValidAdditiveNumber(first, second, index + len(next), length, s)
}
