package solutions

import (
    "sort"
)

func frequencySort(s string) string {
    dictionary := make([]int, 128)
    list := make([]byte, len(s))

    for i := 0; i < len(s); i++ {
        dictionary[s[i]]++
        list[i] = s[i]
    }

    sort.Slice(list, func (i int, j int) bool {
        if dictionary[list[i]] == dictionary[list[j]] {
            return list[i] < list[j]
        }

        return dictionary[list[i]] > dictionary[list[j]]
    })

    return string(list)
}
