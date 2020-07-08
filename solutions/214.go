package solutions

import (
    "bytes"
)

func shortestPalindrome(s string) string {
    length := len(s)

    if length == 0 {
        return ""
    }

    if length == 1 {
        return "a"
    }

    index := computeIndex(s[: length / 2])
    i, j := length - 1, 0

    for i > j {
        if j == -1 || s[i] == s[j] {
            i--
            j++
        } else {
            j = index[j]
        }
    }

    var palindromeEnd int

    if i == j {
        palindromeEnd = 2 * j + 1
    } else {
        palindromeEnd = 2 * j
    }

    var result bytes.Buffer

    for i = length - 1; i >= palindromeEnd; i-- {
        result.WriteByte(s[i])
    }

    result.WriteString(s)

    return result.String()
}

func computeIndex(part string) []int {
    index := make([]int, len(part))
    index[0] = -1

    i, j := 0, -1

    for i < len(part) - 1 {
        if j == -1 || part[j] == part[i] {
            i++
            j++

            if part[j] == part[i] {
                index[i] = index[j]
            } else {
                index[i] = j
            }
        } else {
            j = index[j]
        }
    }

    return index
}
