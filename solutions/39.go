package solutions

import (
    "sort"
)

func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    result := make([][]int, 0)

    sumCandidates(candidates, target, []int{}, &result)

    return result
}

func sumCandidates(candidates []int, target int, current []int, result *[][]int) {
    for i, value := range candidates {
        if target - value <= 0 {
            if target - value == 0 {
                match := make([]int, len(current) + 1)
                copy(match, current)
                match[len(current)] = value

                *result = append(*result, match)
            }

            break
        }

        sumCandidates(candidates[i:], target - value, append(current, value), result)
    }
}