package solutions

import (
    "sort"
)

func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)

    result, i, j := 0, 0, 0

    for i < len(g) && j < len(s) {
        if g[i] <= s[j] {
            result++
            i++
            j++
        } else {
            j++
        }
    }

    return result
}
