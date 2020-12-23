package solutions

import (
    "strings"
)

func repeatedSubstringPattern(s string) bool {
    if len(s) == 0 {
        return false
    }

    size := len(s)
    substring := (s + s)[1: size * 2 - 1]

    return strings.Contains(substring, s)
}
