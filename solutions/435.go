package solutions

import (
    "sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
    result, previous := 0, 0

    sort.Slice(intervals, func(a, b int) bool { return intervals[a][0] < intervals[b][0] })

    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] < intervals[previous][1] {
            if intervals[i][1] < intervals[previous][1] { previous = i }
            result++
        } else {
            previous = i
        }
    }

    return result
}
