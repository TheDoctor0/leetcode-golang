package solutions

import (
    "sort"
)

func findRightInterval(intervals [][]int) []int {
    a := intervals

    for i := 0; i < len(a); i++  {
        a[i] = append(a[i], i)
    }

    sort.Slice(a, func(i int, j int) bool {
        return a[i][0] < a[j][0]
    })

    result := make([]int, len(a))

    for i := 0; i < len(a); i++ {
        b := a[i + 1:]

        position := i + 1 + sort.Search(len(a) - i - 1, func (j int) bool {
            if b[j][0] < a[i][1] {
                return false
            } else {
                return true
            }
        })

        if position >= len(a) {
            result[a[i][2]] = -1
        } else {
            result[a[i][2]] = a[position][2]
        }
    }

    return result
}
