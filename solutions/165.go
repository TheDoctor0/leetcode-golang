package solutions

import (
    "strconv"
    "strings"
)

func compareVersion(version1 string, version2 string) int {
    firstVersion := strings.Split(version1, ".")
    secondVersion := strings.Split(version2, ".")
    length := len(firstVersion)

    if len(secondVersion) > length {
        length = len(secondVersion)
    }

    firstVersionParts := make([]int, length)
    secondVersionParts := make([]int, length)

    for i, rawPart := range firstVersion {
        part, _ := strconv.Atoi(rawPart)

        firstVersionParts[i] = part
    }

    for i, rawPart := range secondVersion {
        part, _ := strconv.Atoi(rawPart)

        secondVersionParts[i] = part
    }

    for i := 0; i < length; i++ {
        if firstVersionParts[i] < secondVersionParts[i] {
            return -1
        } else if firstVersionParts[i] > secondVersionParts[i] {
            return 1
        }
    }

    return 0
}
