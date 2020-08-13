package solutions

import (
    "sort"
)

func hIndex(citations []int) int {
    result := 0

    sort.Ints(citations)

    for i := 0; i < len(citations); i++ {
        if citations[len(citations) - 1 - i] >= i + 1 {
            result = i + 1
        }
    }

    return result
}
