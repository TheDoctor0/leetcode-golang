package solutions

import (
    "sort"
)

func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)

    line := make([]int, 0)
    result := make([][]int, 0)

    calc(0, target, candidates, line, &result)

    return result
}

func calc(index int, target int, candidates []int, line []int, result *[][]int){
    for i := index; i < len(candidates); {
        line = append(line, candidates[i])
        sum := 0

        for _, value := range line {
            sum += value
        }

        if sum == target {
            temp := make([]int, len(line))
            copy(temp, line)
            *result = append(*result, temp)

            return
        }

        if sum > target {
            return
        }

        calc(i + 1, target, candidates, line, result)
        line = line[:len(line) - 1]
        i++

        for ; i < len(candidates) && candidates[i] == candidates[i - 1]; {
            i++
        }
    }
}
