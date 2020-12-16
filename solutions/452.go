package solutions

import (
    "sort"
)

func findMinArrowShots(points [][]int) int {
    if len(points) == 0 {
        return 0
    }

    sort.Slice(points, func(i, j int) bool {
        return points[i][1] < points[j][1]
    })

    arrows := 1
    position := points[0][1]

    for i := 1; i < len(points); i++ {
        if points[i][0] <= position {
            continue
        }

        arrows++
        position = points[i][1]
    }

    return arrows
}
