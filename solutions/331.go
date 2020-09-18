package solutions

import (
    "strings"
)

func isValidSerialization(preorder string) bool {
    if len(preorder) == 0 {
        return false
    }

    nodes := strings.Split(preorder, ",")
    currentLevel, nextLevel, current := 1, 0, 0

    for i, value := range nodes {
        current++

        if value != "#" {
            nextLevel += 2
        }

        if current == currentLevel {
            if nextLevel == 0 && i == len(nodes) - 1 {
                return true
            }

            currentLevel = nextLevel
            nextLevel = 0
            current = 0
        }
    }

    return false
}
