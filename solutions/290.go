package solutions

import (
    "strings"
)

func wordPattern(pattern string, str string) bool {
    list := strings.Split(str, " ")
    firstDict, secondDict := make(map[byte]string), make(map[string]byte)

    if len(pattern) != len(list) {
        return false
    }

    for i := 0; i < len(pattern); i++ {
        if value, ok := firstDict[pattern[i]]; !ok {
            firstDict[pattern[i]] = list[i]
        } else {
            if value != list[i] {
                return false
            }
        }

        if word, ok := secondDict[list[i]]; !ok {
            secondDict[list[i]] = pattern[i]
        } else {
            if word != pattern[i] {
                return false
            }
        }
    }

    return true
}
