package solutions

import (
    "strings"
)

func lengthLongestPath(input string) int {
    paths := strings.Split(input, "\n")
    levelMap := map[int]int{}
    levelMap[0], levelMap[-1] = -1, -2
    current := 0

    for _, path := range paths {
        level, length := getLevel(path)
        levelMap[level] = levelMap[level - 1] + length + 1

        if isFile(path) {
            if levelMap[level] > current {
                current = levelMap[level]
            }
        }
    }

    return current
}

func getLevel(s string) (int, int) {
    for index, value := range s {
        if value != '\t' {
            return index + 1, len(s) - index
        }
    }

    return 0, 0
}

func isFile(s string) bool {
    return strings.Contains(s, ".")
}
