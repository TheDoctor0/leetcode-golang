package solutions

import (
    "sort"
)

func merge(intervals [][]int) [][]int {
    var result [][]int

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    for i, interval := range intervals {
        if i == 0 {
            result = append(result, interval)

            continue
        }

        if len(result) == 0 {
            break
        }

        end := result[len(result) - 1]

        if interval[0] > end[1] {
            result = append(result, interval)
        } else if interval[1] > end[1] {
            end[1] = interval[1]
        }
    }

    return result
}
