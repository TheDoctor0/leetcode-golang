package solutions

import (
    "sort"
)

func reconstructQueue(people [][]int) [][]int {
    sort.Slice(people, func (i int, j int) bool {
        if people[i][0] == people[j][0] {
            return people[i][1] < people[j][1]
        }

        return people[i][0] > people[j][0]
    })

    result := make([][]int, len(people))

    for i := 0; i < len(people); i++ {
        result[i] = people[i]

        for j := people[i][1]; j < i; j++ {
            result[j], result[i] = result[i], result[j]
        }
    }

    return result
}
