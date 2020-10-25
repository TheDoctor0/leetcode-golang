package solutions

import (
    "strconv"
    "strings"
)

func matchingBracket(s string, start int) int {
    count := 0

    for i := start; i < len(s); i++ {
        if s[i] == '[' {
            count++
        } else if s[i] == ']' {
            count--
        }

        if count == 0 {
            return i
        }
    }

    return -1
}

func deserialize(s string) *NestedInteger {
    result := &NestedInteger{}

    if len(s) == 0 {
        return result
    }

    if s[0] == '[' {
        s = s[1: len(s) - 1]
        i := 0

        for i < len(s) {
            if s[i] != '[' {
                part := strings.Index(s[i :], ",")

                if part == - 1{
                    result.Add(*deserialize(s[i:]))
                    i = len(s)
                } else {
                    result.Add(*deserialize(s[i: i + part]))
                    i = i + part + 1
                }
            } else {
                part := matchingBracket(s, i)
                result.Add(*deserialize(s[i: part + 1]))
                i = part + 2
            }
        }
    } else {
        number, _ := strconv.Atoi(s)
        result.SetInteger(number)
    }

    return result
}
